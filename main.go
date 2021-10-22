package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2"
)

func main() {
  pageNum := 1

  for pageNum > 0 {
    var initialReq iPropertyData = fetchiPropData("selangor", pageNum, 100)
    _NextPageNum := initialReq.Data.ACSListing.NextPageToken
    _NextPageNumInt, _ := strconv.Atoi(_NextPageNum)

    text := fmt.Sprintf(`nextPageNum: %d`, pageNum)
    fmt.Println(text)

    if (_NextPageNumInt > 0) {
      pageNum = _NextPageNumInt
      _Items := initialReq.Data.ACSListing.Items
      for _, item := range _Items {
        insertDB("scraper", "iprop", item)
      }
    } else {
      pageNum = 0
      fmt.Println("Done")
    }
  }
}

func insertDB(dbName string, collection string, item interface{}) {
  session, _ := mgo.Dial("127.0.0.1")
	c := session.DB(dbName).C(collection)
	
  defer session.Close()

  c.Insert(item)
}

func fetchiPropData(level1 string, pageToken int, pageSize int) iPropertyData {
  query := fmt.Sprintf(`{
    "operationName": null,
    "variables": {
      "filters": {
        "propertyTypes": [
          "AR"
        ],
        "bedroomRange": {},
        "bathroomRange": {},
        "priceRange": {
          "min": null,
          "max": null
        },
        "builtupSizeRange": {
          "min": null,
          "max": null
        },
        "landSizeRange": {},
        "auction": false,
        "transactedIncluded": false,
        "isOwner": false,
        "hasThreeSixtyImages": false,
        "hasGreatDeal": false,
        "featuredTypes": []
      },
      "channels": [
        "sale",
        "new"
      ],
      "places": [
        {
          "level1": "%s"
        }
      ],
      "pageToken": "%d",
      "pageSize": %d,
      "sortBy": null,
      "customTexts": [],
      "placeIds": [
        "7eacc6298d7c46b29f38fcc5e97a97ea"
      ],
      "developerIds": [],
      "poiIds": [],
      "distance": false,
      "tiers": [],
      "matchCustomTextToPlaceId": false
    },
    "query": "query ($channels: [String!], $customTexts: [String], $placeIds: [String], $poiIds: [String], $distance: Int, $sortBy: String, $filters: ListingFilter, $places: [PlaceFilter], $developerId: String, $developerIds: [String], $pageToken: String, $pageSize: Int, $adId: String, $tiers: [String], $excludeListing: String, $matchCustomTextToPlaceId: Boolean) {\n  ascListings(channels: $channels, customTexts: $customTexts, placeIds: $placeIds, poiIds: $poiIds, distance: $distance, sortBy: $sortBy, filters: $filters, places: $places, pageSize: $pageSize, pageToken: $pageToken, developerId: $developerId, developerIds: $developerIds, primaryListingFromSolr: true, adId: $adId, tiers: $tiers, excludeListing: $excludeListing, matchCustomTextToPlaceId: $matchCustomTextToPlaceId) {\n    items {\n      id\n      channels\n      logo {\n        type\n        url\n      }\n      kind\n      shareLink\n      title\n      description\n      showcase {\n        remark\n      }\n      subtitle\n      tier\n      isPremiumPlus\n      propertyType\n      color\n      prices {\n        type\n        currency\n        label\n        symbol\n        min\n        max\n        minPricePerSizeUnitByBuiltUp\n        maxPricePerSizeUnitByBuiltUp\n        minPricePerSizeUnitByLandArea\n        maxPricePerSizeUnitByLandArea\n        monthlyPayment\n      }\n      cover {\n        type\n        url\n        urlTemplate\n        width\n        height\n        description\n        thumbnailUrl\n        mimeType\n      }\n      medias {\n        type\n        url\n        urlTemplate\n        width\n        height\n        description\n        thumbnailUrl\n        mimeType\n      }\n      unitFeatures {\n        code\n        description {\n          enGB\n          msMY\n        }\n      }\n      buildingFacilities {\n        code\n        description {\n          enGB\n          msMY\n        }\n      }\n      floorPlanImages {\n        type\n        url\n        thumbnailUrl\n        mimeType\n        urlTemplate\n      }\n      youtubeIds\n      updatedAt\n      postedAt\n      address {\n        formattedAddress\n        lat\n        lng\n        hideMarker\n      }\n      referenceCode\n      listerReferenceCode\n      transacted\n      multilanguagePlace {\n        enGB {\n          level1\n          level2\n          level3\n        }\n        msMY {\n          level1\n          level2\n          level3\n        }\n      }\n      organisations {\n        id\n        type\n        name\n        email\n        license\n        website\n        zendeskId\n        award {\n          media {\n            type\n            url\n            urlTemplate\n            width\n            height\n            description\n            thumbnailUrl\n            mimeType\n          }\n          categories\n          status\n          url\n          type\n        }\n        description\n        estimateListsSize {\n          sale\n          rent\n          new\n        }\n        logo {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        color\n        address {\n          formattedAddress\n          lat\n          lng\n          hideMarker\n        }\n        contact {\n          phones {\n            number\n            label\n          }\n          emails\n          bbms\n        }\n        image {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n      }\n      active\n      attributes {\n        bedroom\n        bathroom\n        landArea\n        builtUp\n        carPark\n        rate\n        furnishing\n        floorZone\n        governmentRates\n        buildingAge\n        outsideArea\n        maintenanceFee\n        maintenanceFeeByPsf\n        layout\n        landTitleType\n        tenure\n        isBumiLot\n        topYear\n        aircond\n        pricePSF\n        pricePerSizeUnit\n        pricePerSizeUnitLandArea\n        pricePerSizeUnitBuiltUp\n        minimumPricePerSizeUnit\n        maximumPricePerSizeUnit\n        facingDirection\n        unitType\n        occupancy\n        titleType\n        promotion\n        highlight\n        sizeUnit\n        sizeUnitLandArea\n        auctionDate\n        featureLabel\n        completionStatus\n        projectStage\n        bumiDiscount\n        totalUnits\n        completionDate\n        availableUnits\n        downloadUrl\n        agencyAdvertisingAwardSeal\n        agentAdvertisingAwardSeal\n        developerAdvertisingAwardSeal\n        youtubeId\n        threeDUrl\n        image360\n        ascImage360 {\n          roomType\n          url\n        }\n        hasImage360\n        developerName\n        buildYear\n        projectLicense\n        projectLicenseValidity\n        projectAdvertisingPermit\n        projectAdvertisingPermitValidity\n        projectBuildingReferenceNo\n        projectApprovalAuthorityBuildingPlan\n        projectLandEncumbrance\n        views\n        electricity\n        certificate\n        propertyCondition\n        phoneLine\n        maidRooms\n        maidBathroom\n        ensuite\n        roomType\n        builtYear\n        totalBlocks\n        totalFloors\n        buildingManagement\n        floorHeight\n        characteristicDescription\n        transportDescription\n        governmentWebsite\n        minimumStay\n        architectName\n        contractorName\n        projectType\n        budgetRange\n        buildingId\n        isLinked\n      }\n      greatDeal {\n        title\n        remarks\n        dealEndDate\n        dealStartDate\n        dealTypes {\n          valueCode\n          description\n          displayOrder\n        }\n      }\n      listers {\n        id\n        type\n        name\n        jobTitle\n        knownLanguages\n        license\n        website\n        award {\n          media {\n            type\n            url\n            urlTemplate\n            width\n            height\n            description\n            thumbnailUrl\n            mimeType\n          }\n          categories\n          status\n          url\n          type\n        }\n        description\n        specificPlace\n        estimateListsSize {\n          sale\n          rent\n          new\n        }\n        image {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        color\n        address {\n          formattedAddress\n          lat\n          lng\n          hideMarker\n        }\n        contact {\n          phones {\n            number\n            label\n            via\n          }\n          emails\n          bbms\n        }\n        createdAt\n      }\n      banner {\n        title\n        imageUrl\n        link\n        trackingLink\n        largeImage {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        smallImage {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n      }\n      bankList {\n        data {\n          bank {\n            logo\n            name\n            url\n          }\n          mortgage {\n            interestRate\n            promotionInYear\n            term\n            downPayment\n          }\n        }\n      }\n      rates {\n        builtUp\n        landArea\n        layout\n        sizeUnit\n        price {\n          min\n          max\n          type\n          currency\n        }\n      }\n    }\n    totalCount\n    nextPageToken\n    searchToken\n    multilanguagePlaces {\n      enGB {\n        level1\n        level2\n        level3\n      }\n      msMY {\n        level1\n        level2\n        level3\n      }\n      placeId\n    }\n    locationSpecialists {\n      id\n      type\n      name\n      jobTitle\n      knownLanguages\n      license\n      website\n      award {\n        media {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        categories\n        status\n        url\n        type\n      }\n      description\n      specificPlace\n      estimateListsSize {\n        sale\n        rent\n        new\n      }\n      image {\n        type\n        url\n        urlTemplate\n        width\n        height\n        description\n        thumbnailUrl\n        mimeType\n      }\n      color\n      address {\n        formattedAddress\n        lat\n        lng\n        hideMarker\n      }\n      contact {\n        phones {\n          number\n          label\n        }\n        emails\n        bbms\n      }\n      createdAt\n      organisation {\n        id\n        type\n        name\n        email\n        license\n        website\n        zendeskId\n        award {\n          media {\n            type\n            url\n            urlTemplate\n            width\n            height\n            description\n            thumbnailUrl\n            mimeType\n          }\n          categories\n          status\n          url\n          type\n        }\n        description\n        estimateListsSize {\n          sale\n          rent\n          new\n        }\n        logo {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        color\n        address {\n          formattedAddress\n          lat\n          lng\n          hideMarker\n        }\n        contact {\n          phones {\n            number\n            label\n          }\n          emails\n          bbms\n        }\n        image {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n      }\n    }\n    buildingSpecialists {\n      id\n      type\n      name\n      jobTitle\n      knownLanguages\n      license\n      website\n      award {\n        media {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        categories\n        status\n        url\n        type\n      }\n      description\n      specificPlace\n      estimateListsSize {\n        sale\n        rent\n        new\n      }\n      image {\n        type\n        url\n        urlTemplate\n        width\n        height\n        description\n        thumbnailUrl\n        mimeType\n      }\n      color\n      address {\n        formattedAddress\n        lat\n        lng\n        hideMarker\n      }\n      contact {\n        phones {\n          number\n          label\n        }\n        emails\n        bbms\n      }\n      createdAt\n      organisation {\n        id\n        type\n        name\n        email\n        license\n        website\n        zendeskId\n        award {\n          media {\n            type\n            url\n            urlTemplate\n            width\n            height\n            description\n            thumbnailUrl\n            mimeType\n          }\n          categories\n          status\n          url\n          type\n        }\n        description\n        estimateListsSize {\n          sale\n          rent\n          new\n        }\n        logo {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n        color\n        address {\n          formattedAddress\n          lat\n          lng\n          hideMarker\n        }\n        contact {\n          phones {\n            number\n            label\n          }\n          emails\n          bbms\n        }\n        image {\n          type\n          url\n          urlTemplate\n          width\n          height\n          description\n          thumbnailUrl\n          mimeType\n        }\n      }\n    }\n    administrativeAreas {\n      enGB\n    }\n    poiSuggestions {\n      id\n      type\n      title\n      subtitle\n      label\n      multilanguagePlace {\n        enGB {\n          level1\n          level2\n          level3\n        }\n        msMY {\n          level1\n          level2\n          level3\n        }\n      }\n      additionalInfo {\n        __typename\n        ... on SuggestionAdditionalStationInfo {\n          stops {\n            routeId\n            subType\n            routeColorCode\n            routeDisplaySequence\n            stopIsUnderConstruction\n            stopDisplaySequence\n            routeName {\n              enGB\n              msMY\n            }\n            isExternalStop\n          }\n          name {\n            enGB\n            msMY\n          }\n        }\n      }\n    }\n    developerSuggestions {\n      id\n      multilanguageDeveloper {\n        enGB\n        msMY\n      }\n    }\n  }\n}\n"
  }`, level1, pageToken, pageSize)

  client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://raptor.sg.iproperty.com.my/v1/graphql", strings.NewReader(query))
	req.Header.Add("market", "MY")
	req.Header.Add("x-market", "ipropertymy")
	resp, _ := client.Do(req)

  bodyBytes, _ := ioutil.ReadAll(resp.Body)

  formattedData := iPropertyData{}
  json.Unmarshal(bodyBytes, &formattedData)

  return formattedData
}