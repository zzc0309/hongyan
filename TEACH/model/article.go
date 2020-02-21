package model

type Article struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
	Hits int `json:"hits"`
	//Utime time.Time `json:"utime"`
}
//查询一条
func ArticleGet(id int64)(Article,error){
	mod:=Article{}
	err:=Db.Unsafe().Get(&mod,"select * from article where id=? limit 1",id)
	return mod,err
}
//返回数据列表
func ArticleList()([]Article,error){
	mods:=make([]Article,0,10)
	err:=Db.Unsafe().Select(&mods,"select * from article order by id desc limit 10")
	return mods,err
}
//删除数据
func ArticleDel(id int64)bool{
	res,_:=Db.Exec("delete from article where id=?",id)
	if res==nil{
		return false
	}
	rows,_:=res.RowsAffected()
	if rows>=1{
		return true
	}
	return false
}

func ArcticleAdd(mod *Article)error{
	_,err:=Db.Exec("insert into article (title,author,content,hits) values (?,?,?,?)",mod.Title,mod.Author,mod.Content,mod.Hits)
	return err
}

func ArcticleEdit(mod *Article)error{
	_,err:=Db.Exec("update article set title=?,author=?,content=?,hits=? where id=?",mod.Title,mod.Author,mod.Content,mod.Hits,mod.Id)
	return err
}
