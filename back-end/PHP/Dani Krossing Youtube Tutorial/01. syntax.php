<?php 
// Here we write the PHP code.
?>

<body>
    <p>PHP code can be embedded within HTML.</p>
    <?php echo 'This is a line of PHP code embedded in HTML.'; ?>
</body>

// Use semicolons to terminate statements in PHP
// echo is used to output strings
<?php echo 'This is a statement terminated with a semicolon.'; ?>


// Example #1 PHP Opening and Closing Tags

<?php echo 'if you want to serve PHP code in XHTML or XML documents,
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