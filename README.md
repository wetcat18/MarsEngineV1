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
That will be enough. Now add texture and position to the object creation function. Everything should work. It's not that hard.

# A few more words

I'm not an expert or anything like that. I'm just an amateur, and I'm sure that anyone who really thinks in good architectures can call my code garbage. I'll be glad to have an adequate critic, and if my shit is useful to someone. By the way, in the repository in the file doc.txt the functions of the engine and how to use everything are described in detail. True, this file is written in Russian, but the translator has not been canceled.
