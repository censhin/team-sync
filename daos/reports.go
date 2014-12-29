package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    mod "../models"
    )

func GetReports() (*mod.Reports, error) {
    reports := []mod.Report{}
    reportsMod := new (mod.Reports)

    response, err := GetList("reports")

    err = response.All(&reports)
    count, err := GetCount("reports")

    reportsMod.Reports = reports
    reportsMod.Total = count

    return reportsMod, err
}

func GetReport(reportId string) (mod.Report, error) {
    report := mod.Report{}

    response, err := GetRec("reports", reportId)
    response.Next(&report)

    return report, err
}

func CreateReport(report mod.Report) (mod.Report, error) {
    session := GetSession()

    response, err := db.Table("reports").Insert(report).RunWrite(session)

    if err != nil {
        log.Panic(err)
    }

    report.Id = response.GeneratedKeys[0]

    return report, err
}

func UpdateReport(report mod.Report) (mod.Report, error) {
    session := GetSession()

    _, err := db.Table("reports").Get(report.Id).Update(report).RunWrite(session)

    return report, err
}

func DeleteReport(reportId string) (error) {
    return DeleteRec("reports", reportId)
}