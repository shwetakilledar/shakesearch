import { useState } from 'react'

const useSearch = () => {
  const [searchValue, setSearchValue] = useState('')
  const [loading, setLoading] = useState(false)
  const [hasSuggestion, setHasSuggestion] = useState(false)
  const [data, setData] = useState()
  const [error, setError] = useState(false)
  const [page, setPage] = useState(1)
  const [totalCount, setTotalCount] = useState()

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLoading(false)
    setData(undefined)
    setHasSuggestion(false)
    setSearchValue(e.target.value)
  }

  const handleSearch = async () => {
    if (searchValue.length < 1) {
      return
    }

    setLoading(true)
    await fetch(`/search?q=${searchValue}&page=${page}`).then((response) => {
      response
        .json()
        .then((result) => {
          if (result.data) {
            if (typeof result.data === 'string') {
              setHasSuggestion(true)
            } else {
              setHasSuggestion(false)
              setTotalCount(result.totalCount)
            }
            setData(result.data)
            setLoading(false)
          }
        })
        .catch((err) => {
          setError(error)
          setLoading(false)
        })
    })
  }

  return {
    searchValue,
    onChange,
    loading,
    handleSearch,
    hasSuggestion,
    data,
    setPage,
    totalCount,
  }
}

export default useSearch
