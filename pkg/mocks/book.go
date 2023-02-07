package mocks

import "github.com/karanpratapsingh/tutorials/go/crud/pkg/models"

var Books = []models.Book{
    {
        Id:     1,
        Title:  "Chamber of Secrets",
        Author: "JK Rowling",
        Desc:   "Harry Potter goes to school for the 2nd year.",
    },
    {
        Id:     2,
        Title:  "The Hobbit",
        Author: "JRR Tolkein",
        Desc:   "Hobbits, wizards, elves, galore.",
    },
    {
        Id:     3,
        Title:  "Enders Game",
        Author: "Orson Scott Card",
        Desc:   "Child sci-fi fantasy, inscets vs humans.",
    },
    {
        Id:     4,
        Title:  "Outliers",
        Author: "Malcom Gladwell",
        Desc:   "A book about successful people who put in 10k hours.",
    },
    {
        Id:     5,
        Title:  "Drunk Tank Pink",
        Author: "Adam Grant",
        Desc:   "A book about Human Psycology and why people do what they do.",
    },
}