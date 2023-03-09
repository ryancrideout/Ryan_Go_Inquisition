# Ryan_Go_Inquisition
A repository where I attempt to learn a bit more about the programming language Go.

## Go Resources:
- [Install Go](https://go.dev/doc/install)
- [Go Wikipedia Page](https://en.wikipedia.org/wiki/Go_(programming_language))

## Useful Tutorial Links:
- [Fireship's Go in 100 seconds](https://www.youtube.com/watch?v=446E-r0rXHI)
- [Fireship's Learn Go in 100 lines](https://fireship.io/lessons/learn-go-in-100-lines/)
- [Tutorial Point's Go Tutorial](https://www.tutorialspoint.com/go/index.htm)
- [GoByExample Tutorials](https://gobyexample.com/)
- [Jake Wright's Learn Go in 12 Minutes](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [Go by Example](https://gobyexample.com/) - Great resource for breaking down terms individually.
- [Gophercises](https://gophercises.com/) - Something to do AFTER you've looked at other tutorials.
- [Coursera Go Courses](https://ca.coursera.org/courses?query=golang)

4. When installing pip packages with Python 3.10, you may encounter the error `AttributeError: module 'collections' has no attribute 'MutableMapping'`. To resolve this error, do the following:
    - a. Switch to an alternate version of python using `sudo update-alternatives --config python3`. Python version 3.6 will work fine.
    - b. Run the following commands - the idea is to use Python 3.6 to install `pip` for Python 3.10. Depending on your setup, some of these steps may be redundant - if so, continue to the next step.
    ```
    git clone https://github.com/pypa/setuptools.git && cd setuptools && sudo python3.10 setup.py install
    sudo apt install python3.10-distutils
    curl -sS https://bootstrap.pypa.io/get-pip.py | python3.10
    ```
    - c. Now you should be able to install additional packages for Python 3.10. For additional context for the trouble shooting, go to [here.](https://stackoverflow.com/questions/69512672/getting-attributeerror-module-collections-has-no-attribute-mutablemapping-w/69573368#69573368)
