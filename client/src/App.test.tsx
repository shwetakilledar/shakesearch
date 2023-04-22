import { render, screen } from '@testing-library/react'
import App from './App'

test('renders learn search page', () => {
  render(<App />)
  const searchPage = screen.getByTestId('search-page')
  expect(searchPage).toBeInTheDocument()
})
