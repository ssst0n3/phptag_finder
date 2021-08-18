package phptag_finder

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"testing"
)

func TestFind(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
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
