English version see below after Chinese version.

If you are looking for Windows software, go to [windows-shortcut](https://github.com/Tzxhy/windows-shortcut).


# MacOS上鼠标翻页软件
在mac上，如果有触摸板，可以很容易多指切换屏幕窗口。但对于Mac Mini，或者使用鼠标的用户，如果你的鼠标没有额外的功能键，或者无法装对应的驱动，那么你只能通过键盘：Ctrl + Left/Right 来切换窗口了。通过本工具，可以直接用鼠标中键来切换窗口：

在按下了鼠标中键后的2000ms内：
1. 如果鼠标向左移动超过屏幕宽度的1/20，那么发送按键：`Ctrl+Left`；
2. 如果鼠标向右移动超过屏幕宽度的1/20，那么发送按键：`Ctrl+Right`；
3. 果鼠标向上移动超过屏幕宽度的1/20，那么发送按键：`Ctrl+Up`；
4. 果鼠标向下移动超过屏幕宽度的1/20，那么发送按键：`Ctrl+Down`；

如果超出了2000ms，那么再移动时，也无效。

你需要配置成这样，或者你修改源码，自己编译：
![image](https://github.com/user-attachments/assets/a6209ad7-698b-4f2b-9b39-82a370d304dd)

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
