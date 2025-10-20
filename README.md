# hank

No matter your level of experience, error messages can still be abstract, hard and painful to look at.

hank is a simple CLI utility that focuses on a single task - making error messages digestible and easy to understand. **You can think of this as just another system utility like `grep`, `ls` or `echo` that is at your fingertips for debugging a program.** The current implementation also only uses local language models, which is great for folks who may be a bit more NDA conscious when wanting AI debugging assistance without blindly throwing logs into LLM services (at the plausible compromise of chat response quality).

Simply add `hank` before compiling/running your program to get a clear, readable explanation of any errors that occur.



https://github.com/user-attachments/assets/7c49e0ea-ed79-4307-a2f0-4080e01b48f7



```
hank python <Your Python Program Here>
----------------------------------------
hank make all
----------------------------------------
hank go run main.go
----------------------------------------
etc etc
```

After installing it, you can even get a feel for how it works by just running an executable command like this:
```
hank ls random_directory_that_does_not_exist_on_my_machine
```

### Who is `hank` for?

Every developer/engineer has their own, unique approach to AI in their workflow. `hank` is perfect for those who want AI debugging help at their fingertips without committing to a full agentic terminal solution like Warp.

### Prerequisites

- You will need to have `LMStudio` installed and be serving a language model from `Developer Mode` for this current prototype to work. For more information, please visit [their website](https://lmstudio.ai/). Specific models I use that work well for me are `gemma-3 12B`, `LFM-1.2B` and `gpt-oss 20B`. Feel free to use whatever works best for you.
    - Ensure your model is `Reachable at http://127.0.0.1:1234`

- To build from source, you must have `golang` installed (`>= 1.24.3`)

### Quick Install

If you have `go >= 1.24.3` then you are in luck! `hank` can be installed by running the following:

```
go install github.com/dillondesilva/hank@main
```

Also ensure you have Go's bin folder added to your path. This can be done by adding the following line to your
`.zshrc`, `.bash_profile`, etc config files and running `source ~/.zshrc <Or bash_profile, bashrc etc>`

### Installation

Currently, installation is only supported via building from source:

```
git clone https://github.com/dillondesilva/hank
cd hank
chmod +x install.sh
./install.sh
```

Improvements to this are coming soon

### Contributing Guide

If you have any ideas or feature requests, please open an Issue in the GitHub Issues page.
