<?php 
// Here we write the PHP code.
?>

<body>
    <p>PHP code can be embedded within HTML.</p>
    <?php echo 'This is a line of PHP code embedded in HTML.'; ?>
</body>

<?php 
// Use semicolons to terminate statements in PHP
// echo is used to output strings
echo 'This is a statement terminated with a semicolon.'; 
?>


<?php 
// Example #1 PHP Opening and Closing Tags

echo 'if you want to serve PHP code in XHTML or XML documents,
            use these tags'; ?>

You can use the short echo tag to <?= 'print this string' ?>.
It's equivalent to <?php echo 'print this string' ?>.

<? echo 'this code is within short tags, but will only work '.
        'if short_open_tag is enabled'; ?>

<?php echo 'if a file is pure PHP code, you can omit the closing tag to '.
            'prevent accidental whitespace or new lines being added '.
            'after the PHP code which could cause "headers already sent" '.
            'errors'; ?>


<?php
// Example #2 PHP Code Only File

echo "Hello world\n";

// ... more code
echo "Last statement\n";

// the script ends here with no PHP closing tag
?>


<?php // Example #3 Embedding PHP in HTML ?>

<p>This is going to be ignored by PHP and displayed by the browser.</p>
<?php echo 'While this is going to be parsed.'; ?>
<p>This will also be ignored by PHP and displayed by the browser.</p>


<?php // Example #4 Advanced escaping using conditions ?>

<?php if ($expression == true): ?>
  This will show if the expression is true.
<?php else: ?>
  Otherwise this will show.
<?php endif; ?>


<?php // Example #5 Example showing the closing tag encompassing the trailing newline ?>

<?php echo "Some text"; ?>
No newline
<?= "But newline now" ?>


<?php // Example #6 Examples of entering and exiting the PHP parser ?>

<?php
    echo "This is a test\n";
?>

<?php echo "This is a test\n" ?>

<?php echo "We omitted the last closing tag\n"; ?>


<?php // Example #7 Comments ?>

<?php
    echo "This is a test\n"; // This is a one-line c++ style comment
    /* This is a multi line comment
       yet another line of comment */
    echo "This is yet another test\n";
    echo "One Final Test\n"; # This is a one-line shell-style comment
?>


<?php // Example #8 One Line Comments in HTML ?>

<h1>This is an <?php # echo 'simple';?> example</h1>
<p>The header above will say 'This is an  example'.</p>