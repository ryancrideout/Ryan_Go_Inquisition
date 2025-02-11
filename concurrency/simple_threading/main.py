import multiprocessing
import time

def count_to_ten(process_id):
    for i in range(1000000):
        # print(f"{process_id}: {i}")
        continue

def main():
    start_time = time.time()
    
    # Create a process pool
    with multiprocessing.Pool(processes=multiprocessing.cpu_count()) as pool:
        # Map the function across all inputs
        pool.map(count_to_ten, range(1000))
    
    # Average run time: 1.66 - 2.17 seconds
    # Total execution time: 1.697824 seconds
    # Total execution time: 1.663327 seconds
    # Total execution time: 1.744008 seconds
    # Total execution time: 1.686353 seconds
    # Total execution time: 2.179318 seconds
    # Total execution time: 1.846494 seconds
    elapsed_time = time.time() - start_time
    print(f"Total execution time: {elapsed_time:.6f} seconds")

if __name__ == "__main__":
    main()