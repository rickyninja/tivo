package tivo

const containerInProgressXML = `
<?xml version="1.0" encoding="utf-8"?>
<TiVoContainer xmlns="http://www.tivo.com/developer/calypso-protocol-1.6/">
  <Details>
    <ContentType>x-tivo-container/tivo-videos</ContentType>
    <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
    <Title>Now Playing</Title>
    <LastChangeDate>0x4787065D</LastChangeDate>
    <TotalItems>41</TotalItems>
    <UniqueId>/NowPlaying</UniqueId>
  </Details>
  <SortOrder>Type,CaptureDate</SortOrder>
  <GlobalSort>Yes</GlobalSort>
  <ItemStart>0</ItemStart>
  <ItemCount>41</ItemCount>
  <Item>
    <Details>
      <ContentType>video/x-tivo-raw-tts</ContentType>
      <SourceFormat>video/x-tivo-raw-tts</SourceFormat>
      <Title>Without a Trace</Title>
      <SourceSize>4027580416</SourceSize>
      <Duration>1256000</Duration>
      <CaptureDate>0x478705DE</CaptureDate>
      <EpisodeTitle>4G</EpisodeTitle>
      <Description>A determined private investigator with a reputation for getting the job done disappears while doing undercover surveillance. Copyright Tribune Media Services, Inc.</Description>
      <SourceChannel>705</SourceChannel>
      <SourceStation>KPIXDT</SourceStation>
      <InProgress>Yes</InProgress>
      <HighDefinition>Yes</HighDefinition>
      <ProgramId>EP5240800133</ProgramId>
      <SeriesId>SH524080</SeriesId>
      <ByteOffset>0</ByteOffset>
    </Details>
    <Links>
      <Content>
        <Url>http://652-0001-8039-dd28:80/download/Without%20a%20Trace.TiVo?Container=%2FNowPlaying&amp;id=157800</Url>
        <ContentType>video/x-tivo-raw-tts</ContentType>
        <Available>No</Available>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:in-progress-recording</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
      <TiVoVideoDetails>
        <Url>https://652-0001-8039-dd28:443/TiVoVideoDetails?id=157800</Url>
        <ContentType>text/xml</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </TiVoVideoDetails>
    </Links>
  </Item>
</TiVoContainer>
`
