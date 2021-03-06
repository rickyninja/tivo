package tivo

const nowPlayingXML = `
<?xml version="1.0" encoding="utf-8"?>
<TiVoContainer xmlns="http://www.tivo.com/developer/calypso-protocol-1.6/">
  <Details>
    <ContentType>x-tivo-container/tivo-videos</ContentType>
    <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
    <Title>Now Playing</Title>
    <LastChangeDate>0x508D80B7</LastChangeDate>
    <TotalItems>20</TotalItems>
    <UniqueId>/NowPlaying</UniqueId>
  </Details>
  <SortOrder>Type,CaptureDate</SortOrder>
  <GlobalSort>Yes</GlobalSort>
  <ItemStart>0</ItemStart>
  <ItemCount>20</ItemCount>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>BECKINSALE, KATE</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>2</TotalItems>
      <LastCaptureDate>0x508D648E</LastCaptureDate>
      <UniqueId>2143710</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2143710</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:wishlist-folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>STEWART, KRISTEN</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>3</TotalItems>
      <LastCaptureDate>0x508CB4C6</LastCaptureDate>
      <UniqueId>2143717</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2143717</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:wishlist-folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>PERLMAN, RON</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>1</TotalItems>
      <LastCaptureDate>0x508C676E</LastCaptureDate>
      <UniqueId>2143740</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2143740</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:wishlist-folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>Two and a Half Men</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>11</TotalItems>
      <LastCaptureDate>0x508C6066</LastCaptureDate>
      <UniqueId>820945</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F820945</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>GELLAR, SARAH MICHELLE</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>2</TotalItems>
      <LastCaptureDate>0x508C2F2E</LastCaptureDate>
      <UniqueId>745043</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F745043</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:wishlist-folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>video/x-tivo-raw-pes</ContentType>
      <SourceFormat>video/x-tivo-raw-pes</SourceFormat>
      <Title>Darkness Falls</Title>
      <SourceSize>3380609024</SourceSize>
      <Duration>7201000</Duration>
      <CaptureDate>0x508BE8DE</CaptureDate>
      <Description>The spirit of a woman who was lynched more than 150 years earlier haunts residents of a New England town. Copyright Tribune Media Services, Inc.</Description>
      <SourceChannel>27</SourceChannel>
      <SourceStation>FXP</SourceStation>
      <HighDefinition>No</HighDefinition>
      <ProgramId>MV1304740000</ProgramId>
      <ByteOffset>0</ByteOffset>
    </Details>
    <Links>
      <Content>
        <Url>http://10.0.0.110:80/download/Darkness%20Falls.TiVo?Container=%2FNowPlaying&amp;id=2149665</Url>
        <ContentType>video/x-tivo-raw-pes</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:expires-soon-recording</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
      <TiVoVideoDetails>
        <Url>https://10.0.0.110:443/TiVoVideoDetails?id=2149665</Url>
        <ContentType>text/xml</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </TiVoVideoDetails>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>Angel</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>5</TotalItems>
      <LastCaptureDate>0x508BCCBE</LastCaptureDate>
      <UniqueId>7162</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F7162</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>WHEDON, JOSS</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>1</TotalItems>
      <LastCaptureDate>0x508BCCBE</LastCaptureDate>
      <UniqueId>985367</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F985367</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:wishlist-folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>video/x-tivo-raw-pes</ContentType>
      <SourceFormat>video/x-tivo-raw-pes</SourceFormat>
      <Title>Gladiator</Title>
      <SourceSize>5066719232</SourceSize>
      <Duration>10799000</Duration>
      <CaptureDate>0x508B947E</CaptureDate>
      <Description>Condemned to arena fights by corrupt Roman leader Commodus, Gen. Maximus seeks revenge for his family's deaths. Copyright Tribune Media Services, Inc.</Description>
      <SourceChannel>26</SourceChannel>
      <SourceStation>TNTP</SourceStation>
      <HighDefinition>No</HighDefinition>
      <ProgramId>MV0833410000</ProgramId>
      <ByteOffset>0</ByteOffset>
    </Details>
    <Links>
      <Content>
        <Url>http://10.0.0.110:80/download/Gladiator.TiVo?Container=%2FNowPlaying&amp;id=2149662</Url>
        <ContentType>video/x-tivo-raw-pes</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:expires-soon-recording</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
      <TiVoVideoDetails>
        <Url>https://10.0.0.110:443/TiVoVideoDetails?id=2149662</Url>
        <ContentType>text/xml</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </TiVoVideoDetails>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>video/x-tivo-raw-pes</ContentType>
      <SourceFormat>video/x-tivo-raw-pes</SourceFormat>
      <Title>Fringe</Title>
      <SourceSize>1694498816</SourceSize>
      <Duration>3599000</Duration>
      <CaptureDate>0x508B4E2E</CaptureDate>
      <EpisodeTitle>The Bullet That Saved the World</EpisodeTitle>
      <Description>While tracking a lead into hostile location, the team encounters Phillip Broyles. Copyright Tribune Media Services, Inc.</Description>
      <SourceChannel>10</SourceChannel>
      <SourceStation>KSAZ</SourceStation>
      <HighDefinition>No</HighDefinition>
      <ProgramId>EP010591030098</ProgramId>
      <SeriesId>SH01059103</SeriesId>
      <EpisodeNumber>504</EpisodeNumber>
      <ByteOffset>0</ByteOffset>
    </Details>
    <Links>
      <Content>
        <Url>http://10.0.0.110:80/download/Fringe.TiVo?Container=%2FNowPlaying&amp;id=2143526</Url>
        <ContentType>video/x-tivo-raw-pes</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:expires-soon-recording</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
      <TiVoVideoDetails>
        <Url>https://10.0.0.110:443/TiVoVideoDetails?id=2143526</Url>
        <ContentType>text/xml</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </TiVoVideoDetails>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>Charmed</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>8</TotalItems>
      <LastCaptureDate>0x508AA56E</LastCaptureDate>
      <UniqueId>6052</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F6052</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>Ink Master</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>4</TotalItems>
      <LastCaptureDate>0x508A18CE</LastCaptureDate>
      <UniqueId>2140621</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2140621</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>video/x-tivo-raw-pes</ContentType>
      <SourceFormat>video/x-tivo-raw-pes</SourceFormat>
      <Title>Person of Interest</Title>
      <SourceSize>1694498816</SourceSize>
      <Duration>3601000</Duration>
      <CaptureDate>0x5089FCEA</CaptureDate>
      <EpisodeTitle>Triggerman</EpisodeTitle>
      <Description>Reese and Finch wonder if they should intervene when they learn that a mob enforcer's life is in danger; Finch turns to an unexpected source for help. Copyright Tribune Media Services, Inc.</Description>
      <SourceChannel>5</SourceChannel>
      <SourceStation>KPHO</SourceStation>
      <HighDefinition>No</HighDefinition>
      <ProgramId>EP014198470027</ProgramId>
      <SeriesId>SH01419847</SeriesId>
      <EpisodeNumber>204</EpisodeNumber>
      <ByteOffset>0</ByteOffset>
    </Details>
    <Links>
      <Content>
        <Url>http://10.0.0.110:80/download/Person%20of%20Interest.TiVo?Container=%2FNowPlaying&amp;id=2143738</Url>
        <ContentType>video/x-tivo-raw-pes</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:save-until-i-delete-recording</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
      <TiVoVideoDetails>
        <Url>https://10.0.0.110:443/TiVoVideoDetails?id=2143738</Url>
        <ContentType>text/xml</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </TiVoVideoDetails>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>American Horror Story: Asylum</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>2</TotalItems>
      <LastCaptureDate>0x5088E45E</LastCaptureDate>
      <UniqueId>2141624</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2141624</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>Sons of Anarchy</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>2</TotalItems>
      <LastCaptureDate>0x5087931A</LastCaptureDate>
      <UniqueId>2130627</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2130627</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>video/x-tivo-raw-pes</ContentType>
      <SourceFormat>video/x-tivo-raw-pes</SourceFormat>
      <Title>Fantastic Four: Rise of the Silver Surfer</Title>
      <SourceSize>4223664128</SourceSize>
      <Duration>9000000</Duration>
      <CaptureDate>0x50860126</CaptureDate>
      <Description>Reed, Susan, Johnny and Ben face an intergalactic messenger who has arrived to prepare Earth for destruction. Based on the comic book characters by Stan Lee and Jack Kirby. Copyright Tribune Media Services, Inc.</Description>
      <SourceChannel>27</SourceChannel>
      <SourceStation>FXP</SourceStation>
      <HighDefinition>No</HighDefinition>
      <ProgramId>MV1934440000</ProgramId>
      <ByteOffset>0</ByteOffset>
    </Details>
    <Links>
      <Content>
        <Url>http://10.0.0.110:80/download/Fantastic%20Four%20Rise%20of%20the%20Silver%20Surfer.TiVo?Container=%2FNowPlaying&amp;id=2129851</Url>
        <ContentType>video/x-tivo-raw-pes</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:expired-recording</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
      <TiVoVideoDetails>
        <Url>https://10.0.0.110:443/TiVoVideoDetails?id=2129851</Url>
        <ContentType>text/xml</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </TiVoVideoDetails>
    </Links>
  </Item>
  <Item>
    <Details>
      <ContentType>x-tivo-container/folder</ContentType>
      <SourceFormat>x-tivo-container/tivo-dvr</SourceFormat>
      <Title>The Walking Dead</Title>
      <LastChangeDate>0x508D80B7</LastChangeDate>
      <TotalItems>2</TotalItems>
      <LastCaptureDate>0x5084C4BE</LastCaptureDate>
      <UniqueId>2128240</UniqueId>
    </Details>
    <Links>
      <Content>
        <Url>https://10.0.0.110:443/TiVoConnect?Command=QueryContainer&amp;Container=%2FNowPlaying%2F2128240</Url>
        <ContentType>x-tivo-container/folder</ContentType>
      </Content>
      <CustomIcon>
        <Url>urn:tivo:image:folder</Url>
        <ContentType>image/*</ContentType>
        <AcceptsParams>No</AcceptsParams>
      </CustomIcon>
    </Links>
  </Item>
</TiVoContainer>
`
