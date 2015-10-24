#[Cue descriptions to ASCII](http://52.27.90.205)

##**Usage information**
* Column name detection is case in-sensitive.
* The order of columns does not matter.
* Columns not described here will be ignored.
* You should clear your existing console cues before attempting to import a
generated show; the generated file will not clear them for you.
* Because most spreadsheet applications do not allow you to export multiple sheets
into a single CSV file, users utilizing multiple sheets will need to
upload each sheet individually.

##**Example spreadsheet**
#####Column format specifications are below.
Cue | Description | Pg | Time | Follow | Link
--- |---           |--- |---     |--- |---
.5  | House Lights | 1  | 3      |    |
1   | House Half   |    | 5      | 10 |
1.5 | Blackout     |    | 3/5F10 |    |
2   | Top of show  | 1  | 5/10   |    |
3   | Spooky look  | 5  | 10\5   |    |
P1  | Front light  |    | 15     |    |
5   | Blackout     | 10 | 3.5 	 |    | 1.5

##**Format Details**
###**Cue**
**Accepted values:** Any number

**Example:** 1, 1.5, 4.22

**Note:** Rows without a Cue entry will be ignored
####**Part cues**
**Optional:** Part cues are defined by P\[part number\] (e.g. P1)

**Note:** Parts are assumed to belong to the closest Cue number preceding them
#####**Part example:**

Cue| _
---|---
1  |
P2 | (Cue 1 Part 2)
P3 | (Cue 1 Part 3)
2  |
P2 | (Cue 2 Part 2)

###**Description**
**Accepted values:** Any text.
Total length must be less than 75 characters. Excess characters will be truncated.

###**Page or pg**
**Accepted values:** Any text.
This gets appended to the cue description. Total length of cue description and Page
must be less than 70 characters. Excess characters will be truncated.

###**Time**
**Accepted values:** [time]; [up]/[down]; [up]\[down]

**Example:** 3, 4/6, 5\10

**Note:** If a single number is specified, it will be used for both up and down times

####**Follow times**
**Optional:** You may also specify cue Follows in this column using the format [time]:
F[follow time]

**Example:** 4F10, 4/5F10

**Caution:** This app will accept and output decimal times (e.g. 4.5)
but only you know if your console supports them. Tread lightly.

###**Link**
**Accepted values:** Any number

**Caution:** This app does not verify that the linked cue actually exists

###**Follow**
**Accepted values:** Any number

**Note:** Follows can also be specified in the Time column.
See Time column for more information

##**Compile and run it yourself**
There are no special instructions for building this project.
Simply run `go build main.go` from the source directory.

By default, the application will listen on port 80.
If you wish to change the default port, you can do so in main.go
