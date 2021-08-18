package phptag_finder

import "testing"

func TestFind(t *testing.T) {
	src := []byte(`<html>aaaa
</html>
abcd
<html>bbb
</html>
<?
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
	Find(src)
}
