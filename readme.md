# Switch-WorkSpace-On-Mac
This software can use the middle button of the mouse to switch workspace on Mac for those mouse not having extra functional button.

After pressed the middle button of the mouse, within 2000ms:
1. If your mouse move left exceeds ${screen width / 20}, then trigger `Ctrl + Left`, which will move one workspace towards left by default;
2. If your mouse move right exceeds ${screen width / 20}, then trigger `Ctrl + Right`, which will move one workspace towards right by default;
3. If your mouse move up exceeds ${screen height / 20}, then trigger `Ctrl + Up` ,which will open Scheduler by default;
4. If your mouse move down exceeds ${screen height / 20}, then trigger `Ctrl + Down`.

Be sure, this software uses AppleScript to finish those tasks above, so needs your permission when first executing it.

# Other
If you need it start up after login, just add start-mouse.sh to your login startup items.

# Help
You may need to configurate your keyboard shortcurs like this(if you want to switch workspace):
![image](https://github.com/user-attachments/assets/a6209ad7-698b-4f2b-9b39-82a370d304dd)
