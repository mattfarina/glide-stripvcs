# glide-stripvcs

This is a plugin for [Glide](https://github.com/Masterminds/glide) that strings VCS data from packages in the `vendor/` folder.

## Usage

1. Install Glide and this application (`go get github.com/mattfarina/glide-stripvcs`).
2. Run `glide up` to generate a `glide.lock` file and get the vendor tree set.
3. Run `glide stripvcs` to remove the VCS data from the packages in the `vendor/` folder.
