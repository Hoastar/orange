/*
@Time : 2020/11/9 下午10:57
@Author : hoastar
@File : dashboard
@Software: GoLand
*/

package service

import (
	"database/sql"
	"github.com/hoastar/orange/global/orm"
)

// 查询周统计数据
func WeeklyStatistics() (statisticsData map[string][]interface{}, err error) {
	var (
		// 最近七天代办工单统计图表
		// 时间
		datetime	string
		// 工单总数
		total		int
		// 已完成
		overs		int
		// 待办工单
		processing	int
		//sql查询语句
		sqlValue	string
		// 值
		rows		*sql.Rows
		)

	sqlValue = `SELECT
	a.click_date,
	ifnull( b.total, 0 ) AS total,
	ifnull( b.overs, 0 ) AS overs,
	ifnull( b.processing, 0 ) AS processing 
FROM
	(
	SELECT curdate() AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 1 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 2 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 3 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 4 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 5 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 6 DAY ) AS click_date ) a
	LEFT JOIN (
        SELECT a1.datetime AS datetime,a1.count AS total,b1.count AS overs,c.count AS processing FROM (
            SELECT date( create_time ) AS datetime, count(*) AS count FROM p_work_order_info 
		    GROUP BY date( create_time )) a1
	    LEFT JOIN (SELECT date( create_time ) AS datetime,count(*) AS count FROM p_work_order_info WHERE is_end = 1 GROUP BY date(create_time )) b1 ON a1.datetime = b1.datetime
	    LEFT JOIN (SELECT date( create_time ) AS datetime,count(*) AS count FROM p_work_order_info WHERE is_end = 0 GROUP BY date(create_time )) c ON a1.datetime = c.datetime) b ON a.click_date = b.datetime 
		
		
	order by a.click_date;SELECT
	a.click_date,
	ifnull( b.total, 0 ) AS total,
	ifnull( b.overs, 0 ) AS overs,
	ifnull( b.processing, 0 ) AS processing 
FROM
	(
	SELECT curdate() AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 1 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 2 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 3 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 4 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 5 DAY ) AS click_date UNION ALL
	SELECT date_sub( curdate(), INTERVAL 6 DAY ) AS click_date ) a
	LEFT JOIN (
        SELECT a1.datetime AS datetime,a1.count AS total,b1.count AS overs,c.count AS processing FROM (
            SELECT date( create_time ) AS datetime, count(*) AS count FROM p_work_order_info 
		    GROUP BY date( create_time )) a1
	    LEFT JOIN (SELECT date( create_time ) AS datetime,count(*) AS count FROM p_work_order_info WHERE is_end = 1 GROUP BY date(create_time )) b1 ON a1.datetime = b1.datetime
	    LEFT JOIN (SELECT date( create_time ) AS datetime,count(*) AS count FROM p_work_order_info WHERE is_end = 0 GROUP BY date(create_time )) c ON a1.datetime = c.datetime) b ON a.click_date = b.datetime
	order by a.click_date;`
	rows, err = orm.Eloquent.Raw(sqlValue).Rows()
	if err != nil {
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	// 空map，等价于 statisticsData := make(map[string][]interface)
	statisticsData = map[string][]interface{}{}
	for rows.Next() {
		err = rows.Scan(&datetime, &total, &overs, &processing)
		if err != nil {
			return
		}

		statisticsData["datetime"] = append(statisticsData["datetime"], datetime[:10])
		statisticsData["total"] = append(statisticsData["total"], total)
		statisticsData["overs"] = append(statisticsData["overs"], overs)
		statisticsData["processing"] = append(statisticsData["processing"], processing)
	}
	return
}

// 查询工单提交排名
func SubmitRanking() (submitRankingData map[string][]interface{}, err error) {
	var (
		userId			int
		username		string
		nickname		string
		rankingCount 	int
		rows			*sql.Rows
	)

	sqlValue := `SELECT creator as userId,
       sys_user.username as username,
       sys_user.nick_name as nickname,
       COUNT(*) as rankingCount
	FROM p_work_order_info
    LEFT JOIN sys_user on sys_user.user_id = p_work_order_info.creator
	GROUP BY p_work_order_info.creator ORDER BY rankingCount limit 6;`

	rows, err = orm.Eloquent.Raw(sqlValue).Rows()
	if err != nil {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	submitRankingData = map[string][]interface{}{}
	for rows.Next() {
		submitRankingData["userId"] = append(submitRankingData["userId"], userId)
		submitRankingData["username"] = append(submitRankingData["username"], username)
		submitRankingData["nickname"] = append(submitRankingData["nickname"], nickname)
		submitRankingData["rankingCount"] = append(submitRankingData["rankingCount"], rankingCount)
	}
	return
}
