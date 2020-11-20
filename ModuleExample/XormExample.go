package ModuleExample

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type labor struct {
	Id         int64
	Name       string `xorm:"varchar(20) notnull 'name'"`
	Age        int    `xorm:"int(3) notnull "`
	Occupation string `xorm:"varchar(30) notnull"`
}

var DBSource string = "ml:0000@tcp(192.168.1.103)/person?charset=utf8mb4"
var DBSource_loal string = "root:0000@/person?charset=utf8mb4"

var pln = fmt.Println

func XormExample() {
	var err error
	engine, err = xorm.NewEngine("mysql", DBSource_loal)
	err = engine.Ping()

	var lb labor
	var lbs []labor

	//Query one record
	queryOnerecord := func() {
		return
		pln("Query one record")
		_, err = engine.Table("labor").Where("ID=?", "1").Get(&lb)
		pln(lb)
	}
	queryOnerecord()

	//Query multiple records
	queryMultipleRecord := func() {
		return
		pln("Query multiple records")
		var pwd []string
		engine.Table("labor").Cols("name").Find(&pwd)
		pln(pwd)
	}
	queryMultipleRecord()

	//Query one record
	queryOnerecord2 := func() {
		return
		pln("Query one record")
		cols := []string{"name", "age"}
		var valuesSlice = make([]interface{}, len(cols))
		engine.Table("labor").Cols(cols...).Get(&valuesSlice)
		pln(string(valuesSlice[1].([]byte)))
	}
	queryOnerecord2()

	//Insert
	Insert1 := func() {
		return
		pln("Insert")
		pln(lb)
		//lb.Id = 0
		lb = labor{Name: "tt",
			Age:        11,
			Occupation: "programmer"}
		i64, _ := engine.Insert(lb)
		pln(i64)
	}
	Insert1()

	//check if one record is exist on table.
	checkIfOnerecordisExist := func() {
		return
		pln("check if one record is exist on table.")
		lb = labor{
			Name: "zl",
		}
		ok, _ := engine.Table("labor").Exist(&lb)
		if ok {
			pln("FOUND.")
		}
		if ok, err = engine.Table("labor").Where("id=?", 2).Exist(); ok {
			pln("id=2 found.")
		}
	}
	checkIfOnerecordisExist()

	//query multiple records
	queryMultipleRecords := func() {
		return
		pln("query multiple records")
		err = engine.Table("labor").Where("name=?", "xorm").And("age >?", 1).Limit(10, 0).Find(&lbs)
		pln(lbs)
	}
	queryMultipleRecords()

	JoinTableResult := func() {
		return
		var ress []struct {
			Id             int    `xorm:"id"`
			Name           string `xorm:"name"`
			Age            int    `xorm:"age"`
			Occupation     string `xorm:"occupation"`
			Id_bak         int    `xorm:"id_bak"`
			Name_bak       string `xorm:"name_bak"`
			Occupation_bak string `xorm:"occupation_bak"`
			Password_bak   string `oxrm:"password_bak"`
		}
		err = engine.Table("labor").Select("labor.ID,labor.name,labor.age,labor.occupation,labor_bak.ID as 'id_bak',labor_bak.name as 'name_bak',labor_bak.occupation as 'occupation_bak',labor_bak.password as 'password_bak'").Join("INNER", "labor_bak", "labor.age=labor_bak.age").Limit(10, 0).Where("labor.age=?", 28).Find(&ress)
		pln(ress)
	}
	JoinTableResult()

	//Iterate
	IterateFunc := func() {
		return
		engine.BufferSize(100).Iterate(&labor{Age: 28}, func(idx int, bean interface{}) error {
			//limit
			//idx is index of result,not id.
			if be, ok := bean.(*labor); ok {
				fmt.Println(be.Id)
			}
			return nil
		})
	}
	IterateFunc()

	//Rows
	RowsFunc := func() {
		return
		rows, _ := engine.Rows(&labor{Age: 28})
		defer rows.Close()
		for rows.Next() {
			be := new(labor)
			rows.Scan(be)
			fmt.Println(be.Name)
		}

	}
	RowsFunc()

	//update
	updateFunc := func() {
		return
		//engine.Where("age=?", 11).Update(&labor{Age: 28})
		var i []int = []int{1, 2, 3}
		engine.In("ID", i).Update(&labor{Age: 88})
		// force update indicated columns by Cols
		engine.Id(1).Cols("name").Update(&labor{Name: "hh", Age: 22})

		// force NOT update indicated columns by Omit
		engine.Id(1).Omit("name").Update(&labor{Name: "hh", Age: 22})

		engine.Id(1).AllCols().Update(&labor{Name: "hh", Age: 22})

	}
	updateFunc()

	//delete one or more records, Delete MUST have condition
	deleteFunc := func() {
		return
		//bean也是条件参数，会和where里面的参数合并
		engine.Where("name=?", "a").Delete(&labor{})
	}
	deleteFunc()

	//count
	countFunc := func() {
		return
		i, _ := engine.Count(&labor{Occupation: "Programmar"})
		pln(i)
	}
	countFunc()
	//Sum
	SumFunc := func() {
		//sum of age with named 'xorm'
		return
		agesfloat64, _ := engine.Sum(&labor{Name: "xorm"}, "age")
		fmt.Println(agesfloat64)
	}
	SumFunc()

	//Query conditions builder
	queryConditionFunc := func() {
		//return
		var lbs []labor
		err = engine.Where(builder.NotIn("occupation", "Friar").And(builder.In("age", 88))).Find(&lbs)
		//engine.Table("labor").Where("age=?", 88).Find(&lbs)
		pln(lbs)
	}
	queryConditionFunc()
}

func backup() {
	engine.Insert(&labor{
		Name:       "xorm",
		Age:        233,
		Occupation: "Golang package",
	})
}
