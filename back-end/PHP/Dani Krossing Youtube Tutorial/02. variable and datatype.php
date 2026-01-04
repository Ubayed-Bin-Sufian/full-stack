<?php 
$myName = "Ubayed Bin Sufian";
echo "My name is " . $myName . "<br>";
echo "Welcome to PHP programming!<br>";
?>

<?php
// Example #1 Valid variable names

$var = 'Bob';
$Var = 'Joe';
echo "$var, $Var". "<br>";      // outputs "Bob, Joe"

$_4site = 'not yet';    // valid; starts with an underscore
$täyte = 'mansikka';    // valid; 'ä' is (Extended) ASCII 228.


// Example #2 Invalid variable names

// $4site = 'not yet';     // invalid; starts with a number


// Example #3 Accessing obscure variable names

${'invalid-name'} = 'bar';
$name = 'invalid-name';
echo ${'invalid-name'}, " ", $$name;

?>