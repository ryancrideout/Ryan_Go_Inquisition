#include <iostream>
#include <thread>
#include <vector>
#include <chrono>
#include <mutex>
#include <condition_variable>

class WaitGroup {
private:
    std::mutex mutex;
    std::condition_variable cv;
    int count = 0;

public:
    void add(int delta) {
        std::unique_lock<std::mutex> lock(mutex);
        count += delta;
    }

    void done() {
        std::unique_lock<std::mutex> lock(mutex);
        count--;
        if (count == 0) {
            cv.notify_all();
        }
    }

    void wait() {
        std::unique_lock<std::mutex> lock(mutex);
        cv.wait(lock, [this] { return count == 0; });
    }
};

void count_to_ten(int thread_id, WaitGroup* wait_group) {
    for (int i = 0; i < 1000000; i++) {
        // std::cout << thread_id << ": " << i << std::endl;
        continue;
    }
    wait_group->done();
}

int main() {
    auto start = std::chrono::high_resolution_clock::now();
    
    WaitGroup wait_group;
    std::vector<std::thread> threads;
    
    // Create and start threads
    for (int i = 0; i < 1000; i++) {
        wait_group.add(1);
        threads.emplace_back(count_to_ten, i, &wait_group);
    }

    // Wait for all threads to complete
    wait_group.wait();

    // Join all threads
    for (auto& thread : threads) {
        thread.join();
    }

    auto end = std::chrono::high_resolution_clock::now();
    auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(end - start);

    // Average run time: 0.076 - 0.149 seconds
	// Total execution time: 149ms
	// Total execution time: 104ms
	// Total execution time: 95ms
	// Total execution time: 76ms
	// Total execution time: 81ms
	// Total execution time: 88ms
    std::cout << "Total execution time: " << duration.count() << "ms" << std::endl;

    return 0;
}