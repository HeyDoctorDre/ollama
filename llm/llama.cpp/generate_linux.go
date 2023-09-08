package llm

//go:generate -command cuda-version sh -c "nvcc --version | grep -o -E 'release ([0-9]+)\.' | awk '{print $$2}' | tr -d '.'"
//go:generate cuda-version

//go:generate git submodule init

//go:generate git submodule update --force ggml
//go:generate -command git-apply git -C ggml apply
//go:generate git-apply ../ggml_patch/0001-add-detokenize-endpoint.patch
//go:generate git-apply ../ggml_patch/0002-34B-model-support.patch

//go:generate git submodule update --force gguf

//go:generate sh generate_linux.sh
