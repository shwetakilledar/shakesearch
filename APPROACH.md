# Goal

1. We will be primarily evaluating based on how well the search works for users.
2. A search result with a lot of features (i.e. multi-words and mis-spellings handled), but with results that are hard to read would not be a strong submission.

# Approach

Improved ShakeSearch Relase Date: 23rd April 2023

I prioritised on improving search results by handling multi word search, providing closest matching word for misspelled word from the CompleteWorks and providing case insensitive search.

I converted frontend into React for better and faster UI expeinece by highlting searched word, adding pagination and showing 20 results per page, making ui responsive, and adding skeleton loader for search results for displaying while client is waiting for response from server.

# Demo

https://www.loom.com/share/6ea6791f38814c6ab40b3719ee27428f

# Start

Run following commands to run application in root directory

- `docker build --tag search .`
- `docker run -p 3001:3001 search`

To run Jest test

- `cd shakesearch`
- `npm run test`

To run go tests

- `cd server`
- `go test`

# How to test

1. Go to https://sk-shakesearch.onrender.com/ or run above commands to test the application locally
2. Search for any word from CompleteWorks.txt (e.g `hamlet`)
3. Seach multi words from CompleteWorks.txt (e.g `prince of`)
4. Misspell any word from from CompleteWorks.txt and look for suggestion(s) (e.g `Hamler`)
5. Test results for case insensitivity (e.g `HamlET`)
6. Test responsiveness of the application

# Approach

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

# Changes I would like to prioritize if had more time

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

- Go being new language for me there was a learning curve but  
  had lot of fun while figuring stuff out
- Deploying on Render.com was bit difficult as I am running go
  for the backend and react on the frontend and only one run
  environment could be specified
  - To overcome this problem created a dockerfile
  - Again, this was a learning curve but got to understand docker even more
