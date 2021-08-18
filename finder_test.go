package phptag_finder

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestFind(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	{
		src := []byte(`<html>aaaa
</html>
abcd
<html>bbb
</html>
<?php
echo "Hello world";
?>
<html>bbbb
</html>
<?php
echo "Hello world";
?>
<html>cccc
</html>
`)
		log.Logger.Info(Find(src, "7.4"))
	}
	{
		src, err := ioutil.ReadFile("test/test_data/smarty_internal_template.php")
		assert.NoError(t, err)
		log.Logger.Info(Find(src, "5.6"))
	}
}
