<?php

function count_to_ten($process_id) {
    for ($i = 0; $i < 1000000; $i++) {
        // Commented out to match Go version
        // echo "$process_id: $i\n";
        continue;
    }
}

// Check if pcntl is available
if (!function_exists('pcntl_fork')) {
    die("PCNTL functions not available. Please enable the PCNTL extension.\n");
}

$start = microtime(true);
$children = [];

// Create multiple processes
for ($i = 0; $i < 1000; $i++) {
    $pid = pcntl_fork();
    
    if ($pid == -1) {
        die("Could not fork process\n");
    } else if ($pid === 0) {
        // Child process
        count_to_ten($i);
        exit(0);
    } else {
        // Parent process
        $children[] = $pid;
    }
}

// Wait for all child processes to complete
foreach ($children as $child) {
    pcntl_waitpid($child, $status);
}

// Average run time: 0.68 - 1.23 seconds
// Total execution time: 0.685665 seconds
// Total execution time: 1.105548 seconds
// Total execution time: 1.244796 seconds
// Total execution time: 1.102817 seconds
// Total execution time: 1.234295 seconds
// Total execution time: 0.682322 seconds
$elapsed = microtime(true) - $start;
printf("Total execution time: %.6f seconds\n", $elapsed);

?>
