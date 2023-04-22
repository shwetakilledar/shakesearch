import ReactPaginate from 'react-paginate'

interface SearchResultProps {
  searchValue: string
  data: any
  hasSuggestion: boolean
  loading: boolean
  setPage: (x: number) => void
  handleSearch: () => void
  totalCount: number
}

const SkeletonSearchResult = () => (
  <div className="animate-pulse w-full mt-5">
    <div className="h-2 bg-slate-200 rounded"></div>
  </div>
)

const SearchResult = ({
  searchValue,
  data,
  hasSuggestion,
  loading,
  setPage,
  handleSearch,
  totalCount,
}: SearchResultProps) => {
  const buildResult = (result: string) => {
    const res: any = []
    res.push(
      <tr className="text-justify line-clamp-3">
        {result.split(' ').map((word: string) => (
          <>
            {word.toUpperCase() === searchValue.toUpperCase() ||
            word.includes(searchValue) ? (
              <span className="font-bold">
                {word}
                <span> </span>
              </span>
            ) : (
              <span>
                {word} <span> </span>
              </span>
            )}
          </>
        ))}
      </tr>,
    )
    return res
  }

  return (
    <div className="py-10 px-2.5" data-testid="search-result">
      {loading && (
        <>
          {[...Array(30)].map((_, index) => (
            <SkeletonSearchResult key={index} />
          ))}
        </>
      )}
      {hasSuggestion && (
        <div data-testid="suggestion">
          Did you mean <span className="font-bold">{data}</span>
        </div>
      )}
      {!hasSuggestion && (
        <>
          {data?.length === 0 && <h3 data-testid="no-result">No result</h3>}
          {data?.length > 0 && (
            <h3>
              Total findings:
              <span data-testid="total-findings">{totalCount}</span>
            </h3>
          )}
          {data && (
            <div data-testid="search-results">
              {data.map((result: string, index: number) => (
                <table key={index} className="mt-10">
                  {buildResult(result)}
                </table>
              ))}
              <div className="w-full lg:w-1/2 m-auto pt-10">
                <ReactPaginate
                  breakLabel="..."
                  nextLabel="next >"
                  onPageChange={(e) => {
                    setPage(e.selected)
                    handleSearch()
                  }}
                  pageCount={Math.ceil(totalCount / 20)}
                  previousLabel="< previous"
                  renderOnZeroPageCount={null}
                  className="flex justify-around"
                />
              </div>
            </div>
          )}
        </>
      )}
    </div>
  )
}

export default SearchResult
