# Tasks related to terminal_helper. Improvement points and somethings for v2

1-  Add the logic to see if the file exists at the path and then delete it. - Done
2-  Make all operations using absolute path to keep it generic - Done
3-  Search how can we make a go script run from anywhere regardless of where the
    main.go is.
4-  Add the some UI using bubble tea.
5-  Give some styling using Lip Gloss.
6-  This project should be done end to end no matter what.
7-  Add ps aux commands
8-  Add kill command - Done
9-  Add a way to add tmux themes
10- You can write code to install tmux theme, create a new .tmux file, copy
    contents from old one.
11- Once copied, add the new installation lines in that theme and delete the
    older .tmux files.
12- Add something to write things onto zshell script as well if something
    requires that.

# Update as of 7th August->
1-  Issue encountered in tmux session creation.Using pky package I was able to
    create a tmux session but when it gets created it starts with some garbage
    value plus the tmux has no cursor but still we can write there, not very
    UI appealing.
