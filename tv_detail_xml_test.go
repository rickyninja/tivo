package tivo

// mocked up from tivo xml that was saved as yaml
const tvDetailXML = `
<?xml version="1.0" encoding="utf-8"?>
<TvBusMarshalledStruct:TvBusEnvelope xs:schemaLocation="http://tivo.com/developer/xml/idl/TvBusMarshalledStruct TvBusMarshalledStruct.xsd http://tivo.com/developer/xml/idl/TvPgdRecording TvPgdRecording.xsd http://tivo.com/developer/xml/idl/TvBusDuration TvBusDuration.xsd http://tivo.com/developer/xml/idl/TvPgdShowing TvPgdShowing.xsd http://tivo.com/developer/xml/idl/TvDbShowingBit TvDbShowingBit.xsd http://tivo.com/developer/xml/idl/TvBusDateTime TvBusDateTime.xsd http://tivo.com/developer/xml/idl/TvPgdProgram TvPgdProgram.xsd http://tivo.com/developer/xml/idl/TvDbAdvisory TvDbAdvisory.xsd http://tivo.com/developer/xml/idl/TvDbColorCode TvDbColorCode.xsd http://tivo.com/developer/xml/idl/TvDbMpaa TvDbMpaa.xsd http://tivo.com/developer/xml/idl/TvPgdSeries TvPgdSeries.xsd http://tivo.com/developer/xml/idl/TvDbStar TvDbStar.xsd http://tivo.com/developer/xml/idl/TvPgdChannel TvPgdChannel.xsd http://tivo.com/developer/xml/idl/TvDbBitstreamFormat TvDbBitstreamFormat.xsd" xs:type="TvPgdRecording:TvPgdRecording">
  <recordedDuration>PT1H39M59S</recordedDuration>
  <vActualShowing>
    <element>
    </element>
  </vActualShowing>
  <vBookmark/>
  <showing>
    <showingBits value="4131"/>
    <time>2008-01-08T11:30:00Z</time>
    <duration>PT1H40M</duration>
    <program>
	  <originalAirDate>1998-05-12T01:00:00Z</originalAirDate>
	  <episodeTitle>Becoming</episodeTitle>
	  <episodeNumber>221</episodeNumber>
      <vActor>
        <element>Gellar|Sarah Michelle</element>
        <element>Brendon|Nicholas</element>
        <element>Hannigan|Alyson</element>
        <element>Carpenter|Charisma</element>
        <element>Head|Anthony Stewart</element>
        <element>Boreanaz|David</element>
      </vActor>
      <vChoreographer/>
      <colorCode value="4">COLOR</colorCode>
      <country>United States</country>
      <description>While Angel attempts to unearth a demon who will annihilate the planet, Buffy reluctantly decides to destroy her former lover; Willow unlocks a secret.  Copyright Tribune Media Services, Inc.</description>
      <vDirector>
        <element>Whedon|Joss</element>
      </vDirector>
      <vExecProducer>
        <element>Whedon|Joss</element>
        <element>Berman|Gail</element>
        <element>Gallin|Sandy</element>
        <element>Kuzui|Fran</element>
        <element>Kuzui|Kaz</element>
        <element>Greenwalt|David</element>
      </vExecProducer>
      <vGuestStar/>
      <vHost/>
      <isEpisode>true</isEpisode>
      <vProducer>
        <element>Davies|Gareth</element>
      </vProducer>
      <series>
        <isEpisodic>true</isEpisodic>
        <vSeriesGenre>
          <element>Drama</element>
          <element>Fantasy</element>
          <element>Horror</element>
          <element>Drama</element>
          <element>Sci-Fi and Fantasy</element>
        </vSeriesGenre>
        <seriesTitle>Buffy the Vampire Slayer</seriesTitle>
      </series>
      <title>Buffy the Vampire Slayer</title>
      <vWriter>
        <element>Whedon|Joss</element>
      </vWriter>
    </program>
	<partIndex>1</partIndex>
	<partCount>2</partCount>
    <channel>
      <callsign/>
    </channel>
  </showing>
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
