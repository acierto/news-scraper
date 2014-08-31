describe 'Filter: toGroupFilter', ->
  filter = null

  beforeEach ->
    module 'NewsScraperApp'

  beforeEach ->
    inject (toGroupFilter) ->
      filter = toGroupFilter

  it 'should group articles by group. Scenario 1', ->
    articles = [
      {
        "Link": "http://korrespondent.net/world/russia/3412547-kostroma-desantnyky-na-uchenyiakh-matery-v-nevedenyy-vvs",
        "Title": "СюжетКострома: десантники на \"учениях\", матери в неведении - ВВС",
        "Img": "http://kor.ill.in.ua/m/610x385/1495788.jpg"
      },
      {
        "Link": "http://korrespondent.net/ukraine/3412543-pod-maryupolem-podbyly-kater-ukraynskykh-pohranychnykov",
        "Title": "Под Мариуполем подбили катер украинских пограничников ",
        "Img": ""
      },
      {
        "Link": "http://korrespondent.net/ukraine/events/3412541-uderzhyvaemyi-ukraynskymy-voennymy-aeroport-donetska-okruzhen-opolchentsamy-smy",
        "Title": "Удерживаемый украинскими военными аэропорт Донецка окружен \"ополченцами\" - СМИ",
        "Img": "http://kor.ill.in.ua/m/610x385/1495776.jpg"
      },
      {
        "Link": "http://korrespondent.net/ukraine/3412539-v-snezhnom-separatysty-do-smerty-zamuchyly-voenkoma-mynoborony",
        "Title": "В Снежном сепаратисты до смерти замучили военкома – Минобороны ",
        "Img": ""
      }
    ]

    expect(_.size(filter(articles))).toEqual(2)

  it 'should group articles by group. Scenario 2', ->
    articles = [
      {
        "Link": "http://korrespondent.net/world/russia/3412547-kostroma-desantnyky-na-uchenyiakh-matery-v-nevedenyy-vvs",
        "Title": "СюжетКострома: десантники на \"учениях\", матери в неведении - ВВС",
        "Img": "http://kor.ill.in.ua/m/610x385/1495788.jpg"
      }
    ]

    expect(_.size(filter(articles))).toEqual(1)

  it 'should group articles by group. Scenario 3', ->
    articles = []

    expect(_.size(filter(articles))).toEqual(0)

  it 'should group articles by group. Scenario 4', ->
    articles = [
      {
        "Link": "http://korrespondent.net/world/russia/3412547-kostroma-desantnyky-na-uchenyiakh-matery-v-nevedenyy-vvs",
        "Title": "СюжетКострома: десантники на \"учениях\", матери в неведении - ВВС",
        "Img": "http://kor.ill.in.ua/m/610x385/1495788.jpg"
      },
      {
        "Link": "http://korrespondent.net/ukraine/3412543-pod-maryupolem-podbyly-kater-ukraynskykh-pohranychnykov",
        "Title": "Под Мариуполем подбили катер украинских пограничников ",
        "Img": ""
      },
      {
        "Link": "http://korrespondent.net/ukraine/3412539-v-snezhnom-separatysty-do-smerty-zamuchyly-voenkoma-mynoborony",
        "Title": "В Снежном сепаратисты до смерти замучили военкома – Минобороны ",
        "Img": ""
      }
    ]

    expect(_.size(filter(articles))).toEqual(1)
