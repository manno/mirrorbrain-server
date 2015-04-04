package main

import "io"

const meta4Template = `
<?xml version="1.0" encoding="UTF-8"?>
<metalink xmlns="urn:ietf:params:xml:ns:metalink">
  <generator>MirrorBrain/2.18.1</generator>
  <origin dynamic="true">http://cdn.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4.meta4</origin>
  <published>2015-02-09T23:42:53Z</published>
  <file name="31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">
    <size>577815354</size>

    <!-- <mtime>1419862399</mtime> -->

    <!-- internal id: 31347 -->
    <hash type="md5">769c00c9a6463489762b0eceb5faa678</hash>
    <hash type="sha-1">489f7cff021b19dab0878451d7c0b93fee724ef1</hash>
    <hash type="sha-256">e9b8d1ddb90c2a6ea3a4ca4ddb4c04f85f8f23a88de94f9b5b5648f449ee418a</hash>
    <pieces length="20971520" type="sha-1">
      <hash>357dd8e51f49e54fbbe064dd2527a02db07adafb</hash>
      <hash>7201743e6654b02e3a95c553d14a8c13d2549eb4</hash>
      <hash>3a919734aad8a394cc7647e0486f0a658dc04a93</hash>
      <hash>11134c303a21780071323c8e01977fb17b3dd96e</hash>
      <hash>20dfd04a043edd6656916e3ffc62a30b85d11c64</hash>
      <hash>3d849f8a44bd45ee3546452fe92c257331d45a89</hash>
      <hash>872ec94808442522a64ad7231232ccf4f0969fcb</hash>
      <hash>e102b0d819e6b7975b62accc22839bd9b36a75ef</hash>
      <hash>2dfcfdf721f430ea74b6319d64ecabc9279a6f09</hash>
      <hash>ac0d270617657484498e7efb1559c0f65cc7af88</hash>
      <hash>67aa5266790a9124011ec7897a7168d541cdbe46</hash>
      <hash>6d5b961b9d290dbf1cb1a75bca08b66f1303cbcc</hash>
      <hash>2ceb4f660597ec923defe6573964100208f75fc1</hash>
      <hash>7bd4922a9afa9584fac7a9b092ddfabfd638e84e</hash>
      <hash>9d9a87efc4782a78846919b61d9414b0c565586f</hash>
      <hash>971c5693e9eecc2025d5d451c81d0da27bb331e4</hash>
      <hash>657bad2c55b611bb812ccdb59559d68d62968a13</hash>
      <hash>f7feae32157a2107d7b80166d8b6760c2e8d2f20</hash>
      <hash>f651476c61da277ac2376e7c8ad44e347a7962a1</hash>
      <hash>a68c5a12d419818f4f306cceb776b3beceb62188</hash>
      <hash>e61325e9abf04d65469beaa685e88abe2b720986</hash>
      <hash>f32a98d1d7f2c8c122b33ad8d3bc4820c103d5b9</hash>
      <hash>70344420b3317c7bc158341ff563fbae0ef344eb</hash>
      <hash>90c20b9fa60226c19c8a1f887d263df9d3aaec34</hash>
      <hash>4c398397aab55d6ec4cfe92b1401a7eed289ae3c</hash>
      <hash>dc83265e78b082fd60a36d403beca5a67e7ca411</hash>
      <hash>8239fa062dc79f650f14f8c5c8c931209b3061d2</hash>
      <hash>33ba4c7d56cce8b909641d7a5a2b16601c0e187b</hash>
    </pieces>


    <!-- Found 12 mirrors: 0 in the same network prefix, 1 in the same autonomous system,
         9 handling this country, 1 in the same region, 1 elsewhere -->

    <!-- Mirrors in the same network (78.34.0.0/15): -->

    <!-- Mirrors in the same AS (8422): -->
    <url location="de" priority="1">http://mirror.netcologne.de/CCC/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>

    <!-- Mirrors which handle this country (DE): -->
    <url location="de" priority="2">http://mirror.eu.oneandone.net/projects/media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="3">http://ftp.halifax.rwth-aachen.de/ccc/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="4">http://ftp.chaos-darmstadt.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="5">http://ftp.fau.de/cdn.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="6">http://koeln.ftp.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="7">http://31c3.mirror.speedpartner.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="8">http://ftp.kullen.rwth-aachen.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="9">http://berlin.ftp.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
    <url location="de" priority="10">http://mirror.selfnet.de/CCC/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>

    <!-- Mirrors in the same continent (EU): -->
    <url location="ch" priority="11">http://c3media.vsos.ethz.ch/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>

    <!-- Mirrors in the rest of the world: -->
    <url location="us" priority="12">http://mirror.us.oneandone.net/projects/media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</url>
  </file>
</metalink>
`

// TODO set content type header: "application/metalink4+xml; charset=UTF-8"
func printMeta4(w io.Writer, path string) {
}
