<?php

if ($argc !== 2) {
  echo "Pass in only a filename\n";
  exit;
}

print_r($argv);

$filename = $argv[1];


$s = exec("file ".$filename);
$isJPG = strpos($s, "JPEG") !== FALSE;
if ($isJPG) {
  echo "Yes! This is a jpeg\n";
} else {
  echo "This is not a jpeg\n";
}
