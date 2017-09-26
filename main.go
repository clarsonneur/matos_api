package main

import (
	"database/sql"
	"fmt"
//	"os"
//	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	_ "github.com/mattn/go-adodb"
)

type Materiel struct {
	RefProduit sql.NullString
	TypeProduit sql.NullString
	Marque sql.NullString
	DescriptionProduit sql.NullString
	NoSerie sql.NullString
	Observations sql.NullString
	Absent sql.NullBool
	Rebuter sql.NullBool
	Proprietaire sql.NullString
	Constructeur sql.NullString
	Capacite sql.NullFloat64
	PS sql.NullInt64 //en bars
	PE sql.NullInt64 //en bars
	DatePremEpreuve sql.NullString
	DateDernEpreuve sql.NullString

}

type Materiels map[string]Materiel

func createMdb(f string) error {
	unk, err := oleutil.CreateObject("ADOX.Catalog")
	if err != nil {
		return fmt.Errorf("CreateObject: %s", err)
	}
	cat, err := unk.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return fmt.Errorf("QueryInterface: %s", err)
	}
	_, err = oleutil.CallMethod(cat, "Create", "Provider=Microsoft.Jet.OLEDB.4.0;Data Source=" + f + ";")
	if err != nil {
		return fmt.Errorf("CallMethod: %s", err)
	}
	return nil
}

func main() {
	ole.CoInitialize(0)

	f := "./SAS.mdb"

//	os.Remove(f)

/*	err := createMdb(f)
	if err != nil {
		fmt.Println("create mdb", err)
		return
	}
*/
	db, err := sql.Open("adodb", "Provider=Microsoft.Jet.OLEDB.4.0;Data Source=" + f + ";")
	if err != nil {
		fmt.Println("open", err)
		return
	}

/*	_, err = db.Exec("create table foo (id int not null primary key, name text not null, created datetime not null)")
	if err != nil {
		fmt.Println("create table", err)
		return
	}
*/
/*	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err := tx.Prepare("insert into foo(id, name, created) values(?, ?, ?)")
	if err != nil {
		fmt.Println("insert", err)
		return
	}
	defer stmt.Close()

	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("blabla %d", i), time.Now())
		if err != nil {
			fmt.Println("exec", err)
			return
		}
	}
	tx.Commit()
*/
	matos := make(Materiels)
	rows, err := db.Query("select * from Materiel")
	if err != nil {
		fmt.Println("select", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Materiel
		err = rows.Scan(&p.RefProduit, &p.TypeProduit, &p.Marque, &p.DescriptionProduit, &p.NoSerie, &p.Observations,
			&p.Absent, &p.Rebuter, &p.Proprietaire, &p.Constructeur, &p.Capacite, &p.PS, &p.PE, &p.DatePremEpreuve,
			&p.DateDernEpreuve)
		if err != nil {
			fmt.Println("scan", err)
			return
		}
		matos[p.RefProduit.String] = p
		fmt.Printf("%s - %s\n", p.RefProduit.String, p.TypeProduit.String)
	}
	fmt.Printf("\n%d enregistrements lus.\n", len(matos))
}
