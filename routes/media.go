package routes

import (
    "../models"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-martini/martini"
    "github.com/coopernurse/gorp"
)

func GetMediaIndex(enc Encoder, db gorp.SqlExecutor) (int, string) {
    var media []models.Media
    _, err := db.Select(&media, "select * from media order by id")
    if err != nil {
        checkErr(err, "select failed")
        return http.StatusInternalServerError, ""
    }
    return http.StatusOK, Must(enc.Encode(mediaToIface(media)...))
}

func GetMedia(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Media{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Media)
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddMedia(entity models.Media, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
    err := db.Insert(&entity)
    if err != nil {
        checkErr(err, "insert failed")
        return http.StatusConflict, ""
    }
    w.Header().Set("Location", fmt.Sprintf("/podcaster/media/%d", entity.Id))
    return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateMedia(entity models.Media, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Media{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    oldEntity := obj.(*models.Media)

    entity.Id = oldEntity.Id
    _, err = db.Update(&entity)
    if err != nil {
        checkErr(err, "update failed")
        return http.StatusConflict, ""
    }
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteMedia(db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Media{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Media)
    _, err = db.Delete(entity)
    if err != nil {
        checkErr(err, "delete failed")
        return http.StatusConflict, ""
    }
    return http.StatusNoContent, ""
}

func mediaToIface(v []models.Media) []interface{} {
    if len(v) == 0 {
        return nil
    }
    ifs := make([]interface{}, len(v))
    for i, v := range v {
        ifs[i] = v
    }
    return ifs
}
