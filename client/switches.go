package client

import (
	"flag"
	"fmt"
	"os"
)

// Client interface is the http client for communicating with the server
type Client interface {
	Create(fname, lname string, debit, credit int64) ([]byte, error)
	Edit(id int, fname, lname string, debit, credit int64) ([]byte, error)
	Fetch(id int) ([]byte, error)
	Delete(id int) error
	FetchAll() ([]byte, error)
}

// the CLI command switch
type Switch struct {
	client        Client
	backendAPIURI string
	commands      map[string]func() func(string) error
}

// creates an instance of command Switch
func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s := Switch{client: httpClient, backendAPIURI: uri}
	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit":   s.edit,
		"fetch":  s.fetch,
		"ledger": s.fetchAll,
		"delete": s.delete,
	}
	return s
}

// Executing the command from the command-line is done through this function
func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]
	if !ok {
		return fmt.Errorf("Invalid command '%s'", cmdName)
	}
	return cmd()(cmdName)
}

//prints a useful message about command usage
func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usage of %s:\n<command> [<args>]\n%s", os.Args[0], help)
}

//the create() creates a new transaction in the ledger
func (s Switch) create() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		f, l, d, c := s.ledgerFlags(createCmd)
		if err := s.checkArguments(4); err != nil {
			return err
		}
		if err := s.parseCommand(createCmd); err != nil {
			return err
		}

		res, err := s.client.Create(*f, *l, *d, *c)
		if err != nil {
			return wrapError("Couldn't create the transaction in the ledger", err)
		}
		fmt.Printf("Transaction has been created successfully in the ledger:\n%s", string(res))
		return nil
	}
}

//the edit() edits an existing transaction from the ledger
func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		var id int
		editCmd.IntVar(&id, "id", 0, "The ID (int) of the transaction in the ledger to be edited")
		f, l, d, c := s.ledgerFlags(editCmd)

		if err := s.checkArguments(2); err != nil {
			return err
		}
		if err := s.parseCommand(editCmd); err != nil {
			return err
		}
		if id <= 0 {
			return wrapError("Enter a valid id", nil)
		}

		res, err := s.client.Edit(id, *f, *l, *d, *c)
		if err != nil || string(res) == "" {
			return wrapError("Couldn't edit the transaction in the ledger", err)
		}
		fmt.Printf("Transaction has been edit successfully in the ledger:\n%s", string(res))
		return nil
	}
}

//the fetch() fetches an existing transaction from the ledger
func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		var id int
		fetchCmd.IntVar(&id, "id", 0, "The ID (int) of the transaction in the ledger to be edited")

		if err := s.checkArguments(1); err != nil {
			return err
		}
		if err := s.parseCommand(fetchCmd); err != nil {
			return err
		}
		if id <= 0 {
			return wrapError("Enter a valid id", nil)
		}

		res, err := s.client.Fetch(id)
		if err != nil || string(res) == "" {
			return wrapError("Couldn't fetch the transaction from the ledger", err)
		}
		fmt.Printf("Transaction has been fetched successfully from the ledger:\n%s", string(res))
		return nil
	}
}

//the fetch() fetches an existing transaction from the ledger
func (s Switch) fetchAll() func(string) error {
	return func(cmd string) error {
		fetchAllCmd := flag.NewFlagSet(cmd, flag.ExitOnError)

		if err := s.checkArguments(0); err != nil {
			return err
		}
		if err := s.parseCommand(fetchAllCmd); err != nil {
			return err
		}

		res, err := s.client.FetchAll()
		if err != nil {
			return wrapError("Couldn't fetch the transaction from the ledger", err)
		}
		fmt.Printf("Transaction(s) has been fetched successfully from the ledger:\n%s", string(res))
		return nil
	}
}

//the delete() deletes an existing  transaction from the ledger
func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		deleteCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		var id int
		deleteCmd.IntVar(&id, "id", 0, "The ID (int) of the transaction in the ledger to be edited")

		if err := s.checkArguments(1); err != nil {
			return err
		}
		if err := s.parseCommand(deleteCmd); err != nil {
			return err
		}
		if id <= 0 {
			return wrapError("Enter a valid id", nil)
		}

		err := s.client.Delete(id)
		if err != nil {
			return wrapError("Couldn't delete the transaction from the ledger", err)
		}
		fmt.Printf("Transaction with id:'%d' has been delete successfully from the ledger", id)
		return nil
	}
}

//configres transction specific flags for a command
func (s Switch) ledgerFlags(cmd *flag.FlagSet) (*string, *string, *int64, *int64) {
	f, l, d, c := "", "", int64(0), int64(0)
	cmd.StringVar(&f, "firstname", "", "First Name of the person")
	cmd.StringVar(&f, "f", "", "First Name of the person")
	cmd.StringVar(&l, "lastname", "", "Last Name of the person")
	cmd.StringVar(&l, "l", "", "Last Name of the person")
	cmd.Int64Var(&d, "debit", 0, "Amount owed to the person")
	cmd.Int64Var(&d, "d", 0, "Amount owed to the person")
	cmd.Int64Var(&c, "credit", 0, "Amount owed by the person to you")
	cmd.Int64Var(&c, "c", 0, "Amount owed by the person to you")
	return &f, &l, &d, &c
}

//parses sub-command flags
func (s Switch) parseCommand(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("Couldn't parse '"+cmd.Name()+"' command flags", err)
	}
	return nil
}

// checks if the number of arguments is ok
func (s Switch) checkArguments(minimumArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}
	if len(os.Args)-2 < minimumArgs {
		fmt.Printf("Incorrect use of %s\n%s %s --help\n", os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least: %d arg(s), %d provided", os.Args[1], minimumArgs, len(os.Args)-2)
	}
	return nil
}
