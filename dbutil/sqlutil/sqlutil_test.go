package sqlutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var checkData = "insert into `tb_test` (source,tiny_url,created_on) values (?,?,?) ,[test test test];update `tb_test` set  source=?  where tiny_url=? ,[test value test value];select COALESCE(tiny_url, ''),COALESCE(created_on, ''),COALESCE(origin_url, ''),COALESCE(source, '') from `tb_test` where tiny_url=? and ( created_on=? or source=?) order by created_on asc,tiny_url desc limit 10,[test value or value 2];select count(*) from `tb_test`  where tiny_url=? ,[test value];delete from `tb_test`  where tiny_url=?,[test value];"

func TestExample(t *testing.T) {
	a := exampleBase(false) == checkData
	assert.True(t, a)
}

func TestCOALESCE(t *testing.T) {
	assert.Equal(t, "COALESCE(mini_program_count,'0')", COALESCE("mini_program_count", "0"))
}

func TestExample2(t *testing.T) {
	result := exampleBase(true) != ""
	assert.True(t, result)
}