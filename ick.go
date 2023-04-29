package main

import ("bytes"
         "encoding/binary"
         "fmt"
         "net"
         )

 func ip2Long(ip string) uint32 {
    var long uint32
    binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
    return long
 }

 func ipByte2Long(ip net.IPMask) uint32{
    var long uint32;
    binary.Read(bytes.NewBuffer(ip), binary.BigEndian, &long)
    return long
 }

 func backtoIP4(ipInt uint32) string {

         // need to do two bit shifting and “0xff” masking
         b0 := (ipInt>>24)&0xff
         b1 := (ipInt>>16)&0xff
         b2 := (ipInt>>8)&0xff
         b3 := ipInt & 0xff
         return fmt.Sprintf("%d.%d.%d.%d", b0, b1, b2, b3)
 }

func main() {
    var ip uint32 = ip2Long("8.8.8.8") 
    var mask uint32 = ipByte2Long(net.CIDRMask(21, 32))//0xffff0000
    var revMask uint32 = 0xffffffff^mask
    fmt.Printf("%x %x\n", mask, revMask)
    fmt.Println(ip)
    fmt.Println(uint32(ip&mask))
    var firstIP = backtoIP4(uint32(ip&mask))
    var lastIP = backtoIP4(uint32(ip|revMask))
    var firstUsableIP = backtoIP4(uint32(ip&mask)+1)
    var lastUsableIP = backtoIP4(uint32(ip|revMask)-1)
    fmt.Printf("%s %s %s %s\n", firstIP, firstUsableIP, lastUsableIP, lastIP)
    fmt.Printf("%s\n", backtoIP4(ipByte2Long(net.CIDRMask(21, 32))))
}
