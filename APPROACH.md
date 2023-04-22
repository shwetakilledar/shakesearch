# Goal

1. We will be primarily evaluating based on how well the search works for users.
2. A search result with a lot of features (i.e. multi-words and mis-spellings handled), but with results that are hard to read would not be a strong submission.

# Approach

Improved ShakeSearch Relase Date: 23rd April 2023

I prioritised on improving search results by handling multi word search, providing closest matching word for misspelled word from the CompleteWorks and providing case insensitive search.

I converted frontend into React for better and faster UI expeinece by highlting searched word, adding pagination and showing 20 results per page, making ui responsive, and adding skeleton loader for search results for displaying while client is waiting for response from server.

# Demo

https://www.loom.com/share/6ea6791f38814c6ab40b3719ee27428f

# Run locally

To start the service cd in to project folder and type `go run main.go` in the terminal
Visit http://localhost:3001/

To run Jest test
type `npm run test` in the terminal

To run go tests
cd server
type `go test`

# Priority

Backend

- Handle multi words search
- Suggest closest matching word from the document for mispelled words
  - Suggests closest matching word from the document for mispelled word(s) which have Levenshtein
    Distance of 1
- Test cases
- Case insensitity

Front end

- Convert it into react using typescript
- Responsive UI
- Highlight searched word
- Pagination
- Loader meanwhile fetching the results

Nice to haves I would like to focus on:

- Further optimize search algorithm
- Caching
- Highlight multiple words on ui
- Word suggestion(s)
  - If user has searched for one word
    - Would be nice to have closest matching word(s) as hyper link(s) so
      user doesn't have to type
      the suggested word again in the text box instead just click on it and
      get results from the
      server
  - If user has multiple words and one of the word is misspelled
    - Would be nice to have user click on closest macthing word(s) and
      search the query initially
      entered by user
    - E.g
    - User seraches for `wicke speed`
    - Server gives suggestion(s) for `wicke` as `wicked`
    - When User clicks on suggested word `wicked` it sarches for `wicked speed`

# Difficulties

- Go being new language for me there was a learning curve but had lot of fun while figuring stuff out
