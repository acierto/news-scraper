describe 'Filter: mapByTimeFilter', ->
  filter = null

  beforeEach ->
    module 'NewsScraperApp'

  beforeEach ->
    inject (mapByTimeFilter) ->
      filter = mapByTimeFilter

  it 'should create map with time as a key', ->
    articles = [
      {
        "Source": "http://thekievtimes.ua/",
        "ContentSelector": ".leftCell.mainContentBlock",
        "Link": "http://thekievtimes.ua/economics/397617-slovackij-gaz-tepla-v-kvartirax-ne-budet.html",
        "Title": "Словацкий газ: тепла в квартирах не будет?",
        "Time": "2014-09-04T10:31:00Z"
      }
    ]

    expect(filter(articles)[0].key).toEqual(new Date("2014-09-04T10:31:00Z"))
