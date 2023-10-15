"use strict";

let fs = require('fs');
let data = require('./location-data/aug-05-2023-on-port.json');

const pathData = data.milestones[7].transportRoute.coordinates;
const path = pathData.map((c) => `${c.lng},${c.lat},0`).join(",\n\t\t");
// See https://kml4earth.appspot.com/icons.html for links to icons
// kml: Note that the color is not #rrggbb but #aabbggrr (alpha, blue, green, red)
const kml = `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2">
  <Document>
    <name>Porsche from Emden to Benicia port</name>
    <description>This is path taken by Porsche from port of Emden in Germany to port of Benicia in California.</description>
    <Style id="carmineRedLineGreenPoly">
      <LineStyle>
        <color>ff3800ff</color>
        <width>4</width>
      </LineStyle>
      <PolyStyle>
        <color>ff3800ff</color>
      </PolyStyle>
    </Style>
    <Style id="highlightPlacemark">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/red-stars.png</href>
        </Icon>
      </IconStyle>
    </Style>
    <Style id="yellow-star">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/ylw-stars.png</href>
        </Icon>
      </IconStyle>
    </Style>
    <Style id="yellow-dot">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/ylw-blank-lv.png</href>
        </Icon>
      </IconStyle>
    </Style>    
    <Style id="factoryPlacemark">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/go.png</href>
        </Icon>
      </IconStyle>
    </Style>
    <Style id="emdenPlacemark">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/E.png</href>
        </Icon>
      </IconStyle>
    </Style>
    <Style id="stevensCreekPlacemark">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/S.png</href>
        </Icon>
      </IconStyle>
    </Style>
    <Style id="beniciaPlacemark">
      <IconStyle>
        <Icon>
          <href>https://maps.google.com/mapfiles/kml/paddle/B.png</href>
        </Icon>
      </IconStyle>
    </Style>
    <Placemark>
      <name>The Factory</name>
      <styleUrl>#factoryPlacemark</styleUrl>
      <description>Schwieberdinger Str. 130D, 70435 Stuttgart, Germany. Left on June 26, 2023</description>
      <Point>
        <coordinates>9.152094199999999,48.8350799,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>Shipping route</name>
      <description>Path from Emden to Benicia</description>
      <styleUrl>#carmineRedLineGreenPoly</styleUrl>
      <LineString>
        <extrude>1</extrude>
        <tessellate>1</tessellate>
        <altitudeMode>clampToGround</altitudeMode>
        <coordinates>
        ${path}
        </coordinates>
      </LineString>
    </Placemark>
    <Placemark>
      <name>Port Emden 07-06-2023</name>
      <styleUrl>#emdenPlacemark</styleUrl>
      <description>Start shipment 07-06-2023 </description>
      <Point>
        <coordinates>7.17603,53.3367,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-08-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:19 PM PST</description>
      <Point>
        <coordinates>-9.8939004530628,49.0878753175791,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-09-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:09 PM PST</description>
      <Point>
        <coordinates>-16.0673776197658,48.0362603060919,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-10-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:56 PM PST</description>
      <Point>
        <coordinates>-22.7552892795997,46.4538152055053,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-11-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:19 PM PST</description>
      <Point>
        <coordinates>-28.91101799074779,44.5447241594375,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-12-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:02 PM PST</description>
      <Point>
        <coordinates>-34.719244310746824,42.29774076350066,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-13-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:00 PM PST</description>
      <Point>
        <coordinates>-39.526273810242884,40.0715996182428,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-14-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:21 PM PST</description>
      <Point>
        <coordinates>-44.48591025516438,37.38586948981551,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-15-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:24 PM PST</description>
      <Point>
        <coordinates>-49.39729040565351,34.30167760629351,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-16-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:35 PM PST</description>
      <Point>
        <coordinates>-53.46836487692796,31.400069108333586,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-17-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:16 PM PST</description>
      <Point>
        <coordinates>-57.20851476495191,28.447501796301566,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-18-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:01 PM PST</description>
      <Point>
        <coordinates>-61.05755753731444,25.113154786613975,0</coordinates>
      </Point>
    </Placemark>                             
    <Placemark>
      <name>07-19-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:18 PM PST</description>
      <Point>
        <coordinates>-64.77269395190356,21.61202687947716,0</coordinates>
      </Point>
    </Placemark> 
    <Placemark>
      <name>07-20-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:05 PM PST</description>
      <Point>
        <coordinates>-68.0099077662465,18.421347981370392,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-21-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:48 PM PST</description>
      <Point>
        <coordinates>-72.38172730508852,16.869361677302113,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-22-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:00 PM PST</description>
      <Point>
        <coordinates>-75.90736847797376,13.650581318725092,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-23-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:13 PM PST</description>
      <Point>
        <coordinates>-79.32194937315413,10.369383745478736,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 7:12 AM PST</description>
      <Point>
        <coordinates>-79.5350972444585,8.90524323061693,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 8:20 AM PST</description>
      <Point>
        <coordinates>-79.4596218706555,8.69971704100413,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 9:17 AM PST</description>
      <Point>
        <coordinates>-79.455359272967,8.51661656711562,0</coordinates>
      </Point>
    </Placemark>    
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 10:14 AM PST</description>
      <Point>
        <coordinates>-79.456877469094,8.33073655280564,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 11:33 AM PST</description>
      <Point>
        <coordinates>-79.458934233519,8.07863576721062,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 12:46 PM PST</description>
      <Point>
        <coordinates>-79.4608716528468,7.84087474846628,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 3:46 PM PST</description>
      <Point>
        <coordinates>-79.8025770245288,7.40426764465413,0</coordinates>
      </Point>
    </Placemark> 
    <Placemark>
      <name>07-24-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:59 PM PST</description>
      <Point>
        <coordinates>-81.0796384679874,6.966593503362408,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-25-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:38 PM PST</description>
      <Point>
        <coordinates>-84.74456109342238,8.761687076652033,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-26-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:57 PM PST</description>
      <Point>
        <coordinates>-88.8591670576611,11.09410498163122,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-27-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:59 PM PST</description>
      <Point>
        <coordinates>-93.37136694725237,13.08077084853813,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-28-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:23 PM PST</description>
      <Point>
        <coordinates>-97.58116237871312,15.023660321948055,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-29-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:53 PM PST</description>
      <Point>
        <coordinates>-102.15804930250326,17.240273296968272,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-30-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 9:33 PM PST</description>
      <Point>
        <coordinates>-106.09064061630279,19.536460809861158,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>07-31-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:34 PM PST</description>
      <Point>
        <coordinates>-109.99194786734662,22.085906438639082,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-01-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:40 PM PST</description>
      <Point>
        <coordinates>-113.57397199270802,25.341007038048257,0</coordinates>
      </Point>
    </Placemark>       
    <Placemark>
      <name>08-02-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:15 PM PST</description>
      <Point>
        <coordinates>-116.45580974178779,29.034311694242092,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 10:22 AM PST. 470 miles from Stevens Creek Porsche.</description>
      <Point>
        <coordinates>-118.368534267543,31.2025236052284,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 11:13 AM PST. 459 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-118.486119917703,31.3318102604108,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 1:53 PM PST. 425 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-118.85924251948,31.7390223637388,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 3:39 PM PST. 402 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-119.107019224454,32.0068796213112,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 4:53 PM PST. 386 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-119.283041694907,32.1959285485687,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 5:43 PM PST. 375 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-119.40002171516,32.3209964120613,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 6:47 PM PST. 361 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-119.553062375324,32.4839429060494,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-03-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 8:42 PM PST</description>
      <Point>
        <coordinates>-119.82940104962331,32.77620219387046,0</coordinates>
      </Point>
    </Placemark>    
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 1:35 PM PST. 121 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-121.803011754378,35.5749068148394,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 2:56 PM PST. 105 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-121.938713524327,35.811405168378,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 3:15 PM PST. 100 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-121.972835187764,35.8706275635134,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 4:53 PM PST. 70 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-122.238945433179,36.3291367025546,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 8:10 PM PST. 48 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-122.441359911188,36.7371430114738,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 9:12 PM PST. 42 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-122.551354499139,36.9177136314556,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-04-2023</name>
      <styleUrl>#yellow-star</styleUrl>
      <description>At 10:30 PM PST. 42 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-122.692294453295,37.1378059768877,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>08-05-2023</name>
      <styleUrl>#yellow-dot</styleUrl>
      <description>At 4:54 PM PST. 50 miles from Stevens Creek Porsche</description>
      <Point>
        <coordinates>-122.1554058208113,38.03004493170655,0</coordinates>
      </Point>
    </Placemark>                                           
    <Placemark>
      <name>Port Benicia</name>
      <styleUrl>#beniciaPlacemark</styleUrl>
      <description>Arrived on port at Aug 9 at 5:11 AM PST. California</description>
      <Point>
        <coordinates>-122.135274,38.041906,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>Porsche Stevens Creek</name>
      <styleUrl>#stevensCreekPlacemark</styleUrl>
      <description>4155 Stevens Creek Blvd, Santa Clara, CA 95051. Arrived Aug 15th, 2023 Took possession Aug 17.</description>
      <Point>
        <coordinates>-121.974695,37.3238123,0</coordinates>
      </Point>
    </Placemark>
  </Document>
</kml>`;

fs.writeFileSync('porsche-milestones.kml', kml);
