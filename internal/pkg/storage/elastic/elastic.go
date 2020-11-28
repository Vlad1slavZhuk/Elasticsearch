package elastic

import (
	"Elasticsearch/internal/pkg/constErr"
	"Elasticsearch/internal/pkg/data"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	e "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type StorageElastic struct {
	ESearch  *e.Client
	IndexAcc string
	IndexAds string
}

// AD
type ResultOfSearchAds struct {
	Result ArrHitsAds `json:"hits"`
}

type ArrHitsAds struct {
	Hits []SourceAds `json:"hits"`
}

type SourceAds struct {
	ID     string  `json:"_id"`
	Source data.Ad `json:"_source"`
}

// Account
type ResultOfSearchAccs struct {
	Result ArrHitsAccs `json:"hits"`
}

type ArrHitsAccs struct {
	Hits []SourceAccs `json:"hits"`
}

type SourceAccs struct {
	ID     string       `json:"_id"`
	Source data.Account `json:"_source"`
}

type Count struct {
	Size int `json:"count"`
}

func NewStorageElasticSearch() *StorageElastic {
	es, err := e.NewClient(e.Config{
		Addresses: []string{fmt.Sprintf("http://%v:%v", "elastic", "8005")},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ping check...")
	for _, err := es.Ping(); err != nil; _, err = es.Ping() {
		time.Sleep(2 * time.Second)
	}
	log.Println("Ping success!")
	log.Println(es.Info())

	storage := &StorageElastic{
		ESearch:  es,
		IndexAcc: "accounts",
		IndexAds: "ads",
	}

	return storage
}

func (s *StorageElastic) Add(ad *data.Ad) error {
	size, _ := s.Size()
	ad.ID = uint(size) + 1
	jsonAd, err := json.Marshal(ad)
	if err != nil {
		return err
	}

	res, err := s.ESearch.Index(
		s.IndexAds,              // Index name
		bytes.NewReader(jsonAd), // Document body
		s.ESearch.Index.WithRefresh("true"),
	)

	defer res.Body.Close()

	if err != nil && res.IsError() {
		log.Println("no created")
		return err
	}
	log.Println("created")
	return nil
}

func (s *StorageElastic) Get(id uint) (*data.Ad, error) {
	// creating template query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": id,
			},
		},
	}

	qJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	res, err := s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAds),
		s.ESearch.Search.WithBody(bytes.NewReader(qJSON)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := &ResultOfSearchAds{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return nil, err
	}

	if len(result.Result.Hits) == 0 {
		return nil, constErr.NotFoundAd
	}

	Ad := &result.Result.Hits[0].Source

	return Ad, nil
}

func (s *StorageElastic) GetAll() ([]*data.Ad, error) {
	size, err := s.Size()
	if err != nil || size == 0 {
		return nil, err
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"sort": []map[string]interface{}{ // sorting because we get unsorted data
			map[string]interface{}{
				"id": map[string]interface{}{
					"order": "asc",
				},
			},
		},
	}

	qJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	res, err := s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAds),
		s.ESearch.Search.WithBody(bytes.NewReader(qJSON)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := &ResultOfSearchAds{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return nil, err
	}

	ads := make([]*data.Ad, 0, size)

	for index := range result.Result.Hits {
		ads = append(ads, &result.Result.Hits[index].Source)
	}

	return ads, nil
}

func (s *StorageElastic) Update(temp *data.Ad, id uint) error {
	if _, err := s.Size(); err != nil {
		return err
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": id,
			},
		},
	}

	qJSON, err := json.Marshal(query)
	if err != nil {
		return err
	}

	res, err := s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAds),
		s.ESearch.Search.WithBody(bytes.NewReader(qJSON)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := &ResultOfSearchAds{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return err
	}

	if len(result.Result.Hits) == 0 {
		return constErr.AdBaseIsEmpty
	}

	if result.Result.Hits[0].ID == "" {
		return constErr.AdBaseIsEmpty
	}

	temp.ID = id
	jsAd, err := json.Marshal(temp)
	if err != nil {
		return err
	}

	resUpdate, err := s.ESearch.Index(
		s.IndexAds,
		bytes.NewReader(jsAd),
		s.ESearch.Index.WithDocumentType("_doc"),
		s.ESearch.Index.WithDocumentID(result.Result.Hits[0].ID),
	)
	if err != nil {
		return err
	}
	defer resUpdate.Body.Close()

	return nil
}

func (s *StorageElastic) Delete(id uint) error {
	if _, err := s.Size(); err != nil {
		return err
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": id,
			},
		},
	}

	qJSON, err := json.Marshal(query)
	if err != nil {
		return err
	}

	res, err := s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAds),
		s.ESearch.Search.WithBody(bytes.NewReader(qJSON)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := &ResultOfSearchAds{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return err
	}

	if len(result.Result.Hits) == 0 {
		return constErr.AdBaseIsEmpty
	}

	if result.Result.Hits[0].ID == "" {
		return constErr.AdBaseIsEmpty
	}

	req := esapi.DeleteRequest{
		Index:        s.IndexAds,
		DocumentType: "_doc",
		DocumentID:   result.Result.Hits[0].ID,
		Refresh:      "true",
	}

	resDelete, err := req.Do(context.Background(), s.ESearch)
	if err != nil {
		return err
	}
	defer resDelete.Body.Close()

	query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"sort": []map[string]interface{}{
			map[string]interface{}{
				"id": map[string]interface{}{
					"order": "asc",
				},
			},
		},
	}

	qJSON, err = json.Marshal(query)
	if err != nil {
		return err
	}

	res, err = s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAds),
		s.ESearch.Search.WithBody(bytes.NewReader(qJSON)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	responseByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result = &ResultOfSearchAds{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return err
	}
	// Update id ads
	for index := range result.Result.Hits {
		newID := index + 1
		ad := result.Result.Hits[index].Source
		ad.SetID(uint(newID))

		qJSON, err = json.Marshal(ad)
		if err != nil {
			return err
		}

		reqUpd := esapi.IndexRequest{
			Index:        s.IndexAds,
			DocumentType: "_doc",
			DocumentID:   result.Result.Hits[index].ID,
			Body:         bytes.NewReader(qJSON),
		}

		resUpd, err := reqUpd.Do(context.Background(), s.ESearch)
		if err != nil {
			return err
		}
		defer resUpd.Body.Close()
	}
	return nil
}

func (s *StorageElastic) AddAccount(acc *data.Account) error {
	jsonAcc, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	res, err := s.ESearch.Index(
		s.IndexAcc,               // Index name
		bytes.NewReader(jsonAcc), // Document body
		s.ESearch.Index.WithRefresh("true"),
	)

	defer res.Body.Close()

	if err != nil && res.IsError() {
		log.Println("no created")
		return err
	}
	log.Println("created")
	return nil
}

func (s *StorageElastic) GetAccounts() ([]*data.Account, error) {
	resCount, err := s.ESearch.Count(
		s.ESearch.Count.WithIndex(s.IndexAcc),
	)
	if err != nil {
		return nil, err
	}
	defer resCount.Body.Close()

	count := &Count{}
	if err = json.NewDecoder(resCount.Body).Decode(&count); err != nil {
		return nil, err
	}

	if count.Size == 0 {
		return nil, constErr.AccountBaseIsEmpty
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	qJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	resSearch, err := s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAcc),
		s.ESearch.Search.WithBody(bytes.NewReader(qJSON)),
	)
	if err != nil {
		return nil, err
	}
	defer resSearch.Body.Close()

	responseByte, err := ioutil.ReadAll(resSearch.Body)
	if err != nil {
		return nil, err
	}

	result := &ResultOfSearchAccs{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return nil, err
	}

	accs := make([]*data.Account, 0, count.Size)

	for index := range result.Result.Hits {
		accs = append(accs, &result.Result.Hits[index].Source)
	}

	return accs, nil
}

func (s *StorageElastic) UpdateTokenCurrentAcc(acc *data.Account, token string) error {
	querySearch := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"username": acc.GetUserName(),
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"password": acc.GetPassword(),
						},
					},
				},
			},
		},
	}
	byteQuery, err := json.Marshal(querySearch)
	if err != nil {
		return err
	}

	resSearch, err := s.ESearch.Search(
		s.ESearch.Search.WithIndex(s.IndexAcc),
		s.ESearch.Search.WithBody(bytes.NewReader(byteQuery)),
	)
	defer resSearch.Body.Close()
	if err != nil {
		return err
	}

	responseByte, err := ioutil.ReadAll(resSearch.Body)
	if err != nil {
		return err
	}

	result := &ResultOfSearchAccs{}
	if err := json.Unmarshal(responseByte, result); err != nil {
		return err
	}
	if result.Result.Hits[0].ID == "" {
		return constErr.NotFoundAcc
	}

	acc.SetToken(token)
	jsAcc, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	resUpdate, err := s.ESearch.Index(
		s.IndexAcc,
		bytes.NewReader(jsAcc),
		s.ESearch.Index.WithDocumentType("_doc"),
		s.ESearch.Index.WithDocumentID(result.Result.Hits[0].ID),
	)
	if err != nil {
		return err
	}
	defer resUpdate.Body.Close()

	return nil
}

func (s *StorageElastic) Size() (int, error) {
	res, err := s.ESearch.Count(
		s.ESearch.Count.WithIndex(s.IndexAds),
	)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	count := &Count{}

	if err = json.NewDecoder(res.Body).Decode(&count); err != nil {
		return 0, err
	}

	if count.Size == 0 {
		return 0, constErr.AdBaseIsEmpty
	}

	return count.Size, nil
}
