# Installation MarsEngineV1

For direct installation of the library, you can write a command in the console:

```
go get github.com/wetcat18/MarsEngineV1
```

After installing lib, you can declare it in the code:

```
import "github.com/wetcat18/MarsEngineV1"
```

Wow! You put my engine abstraction on your machine!
It's a small matter, let's figure out how to use it.

# Getting started

If you understand how my code works, then you can use it at your discretion. But if you need everything to work here and now, then here is the architecture I recommend.

```
import engine "github.com/wetcat18/MarsEngineV1"

// The function will be the second thread, you can make changes to objects and build logic in it
func update() {
  // Your updates!
}

// Main func (Wow, that's such a useful comment)
func main() {
	// Window settings
	engine.WindowName("Window")
	engine.WindowSize(1920, 1080)

	// Creating default objects
	engine.CreateObject("just_object", "path_to_texture", 0, 0, true)

	// Launching the update on the second stream
	go update()

	// Starting the engine with the main stream
	engine.Start()
}
```


