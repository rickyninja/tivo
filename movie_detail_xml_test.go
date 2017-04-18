package tivo

const movieDetailXML = `
<?xml version="1.0" encoding="utf-8"?>
<TvBusMarshalledStruct:TvBusEnvelope xs:schemaLocation="http://tivo.com/developer/xml/idl/TvBusMarshalledStruct TvBusMarshalledStruct.xsd http://tivo.com/developer/xml/idl/TvPgdRecording TvPgdRecording.xsd http://tivo.com/developer/xml/idl/TvBusDuration TvBusDuration.xsd http://tivo.com/developer/xml/idl/TvPgdShowing TvPgdShowing.xsd http://tivo.com/developer/xml/idl/TvDbShowingBit TvDbShowingBit.xsd http://tivo.com/developer/xml/idl/TvBusDateTime TvBusDateTime.xsd http://tivo.com/developer/xml/idl/TvPgdProgram TvPgdProgram.xsd http://tivo.com/developer/xml/idl/TvDbAdvisory TvDbAdvisory.xsd http://tivo.com/developer/xml/idl/TvDbColorCode TvDbColorCode.xsd http://tivo.com/developer/xml/idl/TvDbMpaa TvDbMpaa.xsd http://tivo.com/developer/xml/idl/TvPgdSeries TvPgdSeries.xsd http://tivo.com/developer/xml/idl/TvDbStar TvDbStar.xsd http://tivo.com/developer/xml/idl/TvPgdChannel TvPgdChannel.xsd http://tivo.com/developer/xml/idl/TvDbBitstreamFormat TvDbBitstreamFormat.xsd" xs:type="TvPgdRecording:TvPgdRecording">
  <recordedDuration>PT1H39M59S</recordedDuration>
  <vActualShowing>
    <element>
      <showingBits value="4131"/>
      <time>2008-01-08T11:30:00Z</time>
      <duration>PT1H40M</duration>
      <program>
        <vActor>
          <element>Hutcherson|Josh</element>
          <element>Robb|AnnaSophia</element>
          <element>Deschanel|Zooey</element>
          <element>Patrick|Robert</element>
          <element>Clinton|Lauren</element>
          <element>Cerio|Katrina</element>
          <element>Madison|Bailee</element>
          <element>Wood|Devon</element>
          <element>Wakefield|Cameron</element>
          <element>Lawless|Elliott</element>
          <element>Kircher|Isabelle Rose</element>
          <element>Gaines|Latham</element>
          <element>Owen|Carly</element>
        </vActor>
        <vAdvisory>
          <element value="1">LANGUAGE</element>
          <element value="10">ADULT_SITUATIONS</element>
        </vAdvisory>
        <vChoreographer/>
        <colorCode value="4">COLOR</colorCode>
        <country>United States</country>
        <description>A boy and his new friend, the class outsider, create an imaginary world in which they rule as king and queen. Based on the novel by Katherine Paterson. Copyright Tribune Media Services, Inc.</description>
        <vDirector>
          <element>Csupo|Gabor</element>
        </vDirector>
        <vExecProducer/>
        <vProgramGenre>
          <element>Fantasy</element>
        </vProgramGenre>
        <vGuestStar/>
        <vHost/>
        <isEpisode>true</isEpisode>
        <movieRunTime>PT1M36S</movieRunTime>
        <movieYear>2007</movieYear>
        <mpaaRating value="2">PG</mpaaRating>
        <vProducer>
          <element>Lieberman|Hal</element>
          <element>Levine|Lauren</element>
        </vProducer>
        <series>
          <isEpisodic>false</isEpisodic>
          <vSeriesGenre>
            <element>Fantasy</element>
            <element>Movies</element>
            <element>Sci-Fi and Fantasy</element>
          </vSeriesGenre>
          <seriesTitle>Bridge to Terabithia</seriesTitle>
        </series>
        <starRating value="5">THREE</starRating>
        <title>Bridge to Terabithia</title>
        <vWriter>
          <element>Paterson|David</element>
          <element>Stockwell|Jeff</element>
          <element>Wade|Kevin</element>
        </vWriter>
      </program>
    </element>
  </vActualShowing>
  <vBookmark/>
  <showing>
    <showingBits value="4131"/>
    <time>2008-01-08T11:30:00Z</time>
    <duration>PT1H40M</duration>
    <program>
      <vActor>
        <element>Hutcherson|Josh</element>
        <element>Robb|AnnaSophia</element>
        <element>Deschanel|Zooey</element>
        <element>Patrick|Robert</element>
        <element>Clinton|Lauren</element>
        <element>Cerio|Katrina</element>
        <element>Madison|Bailee</element>
        <element>Wood|Devon</element>
        <element>Wakefield|Cameron</element>
        <element>Lawless|Elliott</element>
        <element>Kircher|Isabelle Rose</element>
        <element>Gaines|Latham</element>
        <element>Owen|Carly</element>
      </vActor>
      <vAdvisory>
        <element value="1">LANGUAGE</element>
        <element value="10">ADULT_SITUATIONS</element>
      </vAdvisory>
      <vChoreographer/>
      <colorCode value="4">COLOR</colorCode>
      <country>United States</country>
      <description>A boy and his new friend, the class outsider, create an imaginary world in which they rule as king and queen. Based on the novel by Katherine Paterson. Copyright Tribune Media Services, Inc.</description>
      <vDirector>
        <element>Csupo|Gabor</element>
      </vDirector>
      <vExecProducer/>
      <vProgramGenre>
        <element>Fantasy</element>
      </vProgramGenre>
      <vGuestStar/>
      <vHost/>
      <isEpisode>true</isEpisode>
      <movieRunTime>PT1M36S</movieRunTime>
      <movieYear>2007</movieYear>
      <mpaaRating value="2">PG</mpaaRating>
      <vProducer>
        <element>Lieberman|Hal</element>
        <element>Levine|Lauren</element>
      </vProducer>
      <series>
        <isEpisodic>false</isEpisodic>
        <vSeriesGenre>
          <element>Fantasy</element>
          <element>Movies</element>
          <element>Sci-Fi and Fantasy</element>
        </vSeriesGenre>
        <seriesTitle>Bridge to Terabithia</seriesTitle>
      </series>
      <starRating value="5">THREE</starRating>
      <title>Bridge to Terabithia</title>
      <vWriter>
        <element>Paterson|David</element>
        <element>Stockwell|Jeff</element>
        <element>Wade|Kevin</element>
      </vWriter>
    </program>
    <channel>
      <callsign/>
    </channel>
  </showing>
  <startTime>2008-01-08T11:29:58Z</startTime>
  <stopTime>2008-01-08T13:10:00Z</stopTime>
  <bitstreamFormat>
    <vFormat>
      <element>
        <vByte>
          <base64>EjQAAwABAjoBywxXAAAADwAAAAQAAAACAAAAAwAAAA==</base64>
        </vByte>
      </element>
    </vFormat>
  </bitstreamFormat>
  <expirationTime>2008-01-10T11:30:00Z</expirationTime>
</TvBusMarshalledStruct:TvBusEnvelope>
`
