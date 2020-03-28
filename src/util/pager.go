package util

/**
 * paging tool
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-28
 */
type Pager struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func Paging(count int, pageNo int, pageSize int, list interface{}) Pager {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Pager{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}
