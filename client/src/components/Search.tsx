import useSearch from '../hooks/useSearch'
import SearchResult from './SearchResult'

const Search = () => {
  const {
    onChange,
    searchValue,
    loading,
    handleSearch,
    hasSuggestion,
    data,
    setPage,
    totalCount,
  } = useSearch()

  return (
    <div className="flex flex-col">
      <div className="w-full flex flex-col md:flex-row items-center justify-center m-auto">
        <input
          className="w-full shadow appearance-none border rounded md:w-3/4 py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline !p-5"
          placeholder="Search in The Complete Works of William Shakespeare"
          tabIndex={0}
          defaultValue={searchValue}
          onChange={(e) => {
            onChange(e)
          }}
        />
        <button
          className="w-full p-2.5 md:w-1/4 md:ml-5 md:p-5 text-center border-[1px] rounded-2xl border-gray-500 hover:bg-[#33bd89] hover:border-[#33bd89]"
          tabIndex={0}
          disabled={loading}
          onClick={() => {
            handleSearch()
          }}
          onKeyDown={() => {
            handleSearch()
          }}
        >
          Search
        </button>
      </div>

      <SearchResult
        data={data}
        hasSuggestion={hasSuggestion}
        loading={loading}
        searchValue={searchValue}
        setPage={setPage}
        handleSearch={handleSearch}
        totalCount={totalCount || 0}
      />
    </div>
  )
}

export default Search
