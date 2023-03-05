package main

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadError := godotenv.Load()
	if loadError != nil {
		log.Fatal("Error loading .env file")
	}

	db, openError := badger.Open(badger.DefaultOptions(os.Getenv("BADGER_DB_DATA_PATH")))

	if openError != nil {
		log.Fatal(openError)
	}

	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "badger-db",
			"context": os.Getenv("BADGER_DB_CONTEXT"),
			"version": "0.0.1",
			"database": fiber.Map{
				"storage": "data/badger",
				"version": "v3.2103.5",
			},
		})
	})

	app.Head("/items/:key", func(c *fiber.Ctx) error {
		viewError := db.View(func(txn *badger.Txn) error {
			_, getError := txn.Get([]byte(c.Params("key")))
			if getError != nil {
				return getError
			}

			return nil
		})

		if viewError != nil {
			if viewError == badger.ErrKeyNotFound {
				return c.Status(fiber.StatusNotFound).Send(nil)
			} else {
				return viewError
			}
		}

		return c.Status(fiber.StatusNoContent).Send(nil)
	})

	app.Get("/items/:key", func(c *fiber.Ctx) error {
		var value []byte

		viewError := db.View(func(txn *badger.Txn) error {
			item, getError := txn.Get([]byte(c.Params("key")))
			if getError != nil {
				return getError
			}

			var valueCopyError error
			value, valueCopyError = item.ValueCopy(nil)

			if valueCopyError != nil {
				return valueCopyError
			}

			return nil
		})

		if viewError != nil {
			if viewError == badger.ErrKeyNotFound {
				return c.Status(fiber.StatusNotFound).Send(nil)
			} else {
				return viewError
			}
		}

		return c.SendString(string(value))
	})

	app.Put("/items/:key", func(c *fiber.Ctx) error {
		exists := true

		updateError := db.Update(func(txn *badger.Txn) error {
			_, getError := txn.Get([]byte(c.Params("key")))

			if getError == badger.ErrKeyNotFound {
				exists = false
			} else if getError != nil {
				return getError
			}

			return txn.Set([]byte(c.Params("key")), c.Body())
		})

		if updateError != nil {
			return updateError
		}

		if exists {
			return c.Status(fiber.StatusNoContent).Send(nil)
		} else {
			return c.Status(fiber.StatusCreated).Send(nil)
		}
	})

	app.Delete("/items/:key", func(c *fiber.Ctx) error {
		updateError := db.Update(func(txn *badger.Txn) error {
			_, getError := txn.Get([]byte(c.Params("key")))

			if getError != nil {
				return getError
			}

			deleteError := txn.Delete([]byte(c.Params("key")))

			if deleteError != nil {
				return deleteError
			}

			return nil
		})

		if updateError != nil {
			if updateError == badger.ErrKeyNotFound {
				return c.Status(fiber.StatusNotFound).Send(nil)
			} else {
				return updateError
			}
		}

		return c.Status(fiber.StatusNoContent).Send(nil)
	})

	app.Listen(os.Getenv("BADGER_DB_HOST") + ":" + os.Getenv("BADGER_DB_PORT"))
}
