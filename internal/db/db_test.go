package db

import (
	_ "net/http/pprof"
	"testing"
)

func Test_ReadWriteToBolt(t *testing.T) {

	//t.SkipNow()
	if err := Init("."); err != nil {
		t.Fatal(err)
	}

	dbMsg := DbMessage{
		ID:                "",
		SpecID:            100,
		MsgID:             1,
		RequestTS:         436466364678,
		ResponseTS:        767366436647,
		RequestMsg:        "110100101010010101010",
		ParsedRequestMsg:  nil,
		ResponseMsg:       "11110........",
		ParsedResponseMsg: nil,
		HostAddr:          "localhost:7777",
	}
	for i := 0; i < 10; i++ {
		if err := Write(dbMsg); err != nil {
			t.Fatal(err)
		}

		//time.Sleep(1 * time.Second)
	}

	entries, err := ReadLast(100, 1, 5)
	if entries == nil {
		t.Fatal("No entries found!")
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Log(entries)
}

/*func Test_Read(t *testing.T) {

log.SetLevel(log.DebugLevel)

/*go func() {
	log.Fatal(http.ListenAndServe("localhost:8765", nil))
}()*/

//t.SkipNow()
/*
	if err := Init("."); err != nil {
		t.Fatal(err)
	}

	entries, err := ReadLast(100, 1, 90)
	if entries == nil {
		t.Fatal("No entries found!")
	}
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		t.Log(e)
	}
	t.Log(len(entries))

}*/
