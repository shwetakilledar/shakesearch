import { render, screen } from '@testing-library/react'
import { SearchResult } from '../'

const defaultProps = {
  searchValue: '',
  data: undefined,
  hasSuggestion: false,
  loading: false,
  setPage: (x: number) => {},
  handleSearch: () => {},
  totalCount: 0,
}

test('renders search results component', () => {
  render(<SearchResult {...defaultProps} />)
  const searchResult = screen.getByTestId('search-result')
  expect(searchResult).toBeInTheDocument()
})

test('renders search results with expected result', () => {
  const props = {
    ...defaultProps,
    searchValue: 'nothing',
    totalCount: 5,
    data: ['nothing', 'nothing', 'nothing', 'nothing', 'nothing'],
  }

  render(<SearchResult {...props} />)
  const searchResult = screen.getByTestId('search-results')
  expect(searchResult).toBeInTheDocument()
  const totalFindings = screen.getByTestId('total-findings')
  expect(totalFindings).toBeInTheDocument()
})

test('renders search result component with no result', () => {
  const props = {
    ...defaultProps,
    searchValue: 'something',
    totalCount: 0,
    data: [],
  }

  render(<SearchResult {...props} />)
  const noResult = screen.getByTestId('no-result')
  expect(noResult).toBeInTheDocument()
})

test('renders search result component with suggestion', () => {
  const props = {
    ...defaultProps,
    searchValue: 'somethig',
    data: 'something',
    hasSuggestion: true,
  }

  render(<SearchResult {...props} />)
  const suggestion = screen.getByTestId('suggestion')
  expect(suggestion).toBeInTheDocument()
})

test('renders search result component with multi word suggestions', () => {
  const props = {
    ...defaultProps,
    searchValue: 'hamler',
    data: 'hamlet hammper',
    hasSuggestion: true,
  }

  render(<SearchResult {...props} />)
  const suggestion = screen.getByTestId('suggestion')
  expect(suggestion).toBeInTheDocument()
})
