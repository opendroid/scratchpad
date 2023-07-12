"use strict";

let fs = require('fs');
let data = require('./july-11-2023.json');

const pathData = data.milestones[7].transportRoute.coordinates;
const path = pathData.map((c) => `${c.lng},${c.lat}`).join(",\n\t\t\t\t");
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
          <href>http://maps.google.com/mapfiles/kml/paddle/red-stars.png</href>
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
      <description>Schwieberdinger Str. 130D, 70435 Stuttgart, Germany</description>
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
        <altitudeMode>absolute</altitudeMode>
        <coordinates>
        ${path}
        </coordinates>
      </LineString>
    </Placemark>
    <Placemark>
      <name>Port Emden</name>
      <styleUrl>#emdenPlacemark</styleUrl>
      <description>Start shipment</description>
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
      <name>Port Benicia</name>
      <styleUrl>#beniciaPlacemark</styleUrl>
      <description>California</description>
      <Point>
        <coordinates>-122.135274,38.041906,0</coordinates>
      </Point>
    </Placemark>
    <Placemark>
      <name>Porsche Stevens Creek</name>
      <styleUrl>#beniciaPlacemark</styleUrl>
      <description>4155 Stevens Creek Blvd, Santa Clara, CA 95051</description>
      <Point>
        <coordinates>-121.974695,37.3238123,0</coordinates>
      </Point>
    </Placemark>
  </Document>
</kml>`;
fs.writeFileSync('porsche-2.kml', kml);