curl -XPUT http://localhost:9200/article -d '
{
    "mappings": {
      "properties": {
          "id": {
              "type": "long"
          },
          "user_id": {
              "type": "long"
          },
          "title": {
              "type": "text", "analyzer": "smartcn"
          },
          "summary": {
              "type": "text", "analyzer": "smartcn"
          },
          "content": {
              "type": "text", "analyzer": "smartcn"
          },
          "tags": {
            "type":"keyword"
          },
          "publish_time": {
              "type": "date", "format": "strict_date_optional_time||epoch_second"
          }
        }
    },
    "settings": {
        "index": {
            "sort.field": ["publish_time"],
            "sort.order": ["desc"]
        }
    }
}
'