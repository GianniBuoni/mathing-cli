# ➕ It's MATHEMATICAL!

I often find myself needing to split grociery calcs between me an my roommate.  
And, usually the receipts are very similar to each other.

So I got to thinking: why not just create a SQLite database that stores these receipts  
and then access them with a CLI utility?

## ⚡ Goal

The hope is that I just run the app, load up the last receipt, and then edit  
the quantities and change some of the prices as needed.

Treating this as a fun way to learn how to use the Bubbletea and Lipgloss libraries.

## 🧮 Usage

Invoking `mathing` without any arguments will bring up the receipt TUI.  
The TUI can handle adding, editing, and deleting list items from the receipt and item tables.

Commands:

- `list` : Also brings up the TUI.
- `calc` : Finalizes the calculations. Prints out how much is owed by each user that appeared in the receipt.
- `help` : Prints out a table of available commands.

Development Commands:

- `seed` : Seeds the SQLite database with dummy data.

## ✅ TO DO

- [x] Rework the receipt form/builder to be able to filter through all the available items.
- [ ] Rework how the TUI handles the table data refreshing. Currently, displayed data can become unsynced with db data.
- [ ] Integrate the calc command into the TUI or create a pane of live calcs.
- [ ] Move SQLite db to somewhere in $HOME.
- [x] Reset command to clean out table data / create the option to build a new receipt from scratch.
- [ ] Receipt history. Command to write calc results to a CSV or another table.
