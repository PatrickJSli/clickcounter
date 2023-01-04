# clicking-is-evil

Bash script to display mouse click count in real time. Created this so I could improve at using keyboard shortcuts and see how many times I was using the mouse.

# Usage

1. run `xinput list` to find device id of mouse or touchpad
2. change `12` on line 10 to your device id
3. run the script

# To-do
- accept device id as parameter
- use ascii library like banner to make text larger
- smart repaint to avoid flickering
