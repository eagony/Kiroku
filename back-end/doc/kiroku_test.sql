/*
 Navicat Premium Data Transfer

 Source Server         : 106.14.26.172
 Source Server Type    : MySQL
 Source Server Version : 50648
 Source Host           : 106.14.26.172:3306
 Source Schema         : kiroku_test

 Target Server Type    : MySQL
 Target Server Version : 50648
 File Encoding         : 65001

 Date: 12/10/2020 14:58:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_tags
-- ----------------------------
DROP TABLE IF EXISTS `blog_tags`;
CREATE TABLE `blog_tags`  (
  `blog_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `tag_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`blog_id`, `tag_id`) USING BTREE,
  INDEX `fk_blog_tags_tag`(`tag_id`) USING BTREE,
  CONSTRAINT `fk_blog_tags_blog` FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_blog_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of blog_tags
-- ----------------------------
INSERT INTO `blog_tags` VALUES (3, 7);
INSERT INTO `blog_tags` VALUES (5, 7);
INSERT INTO `blog_tags` VALUES (3, 8);
INSERT INTO `blog_tags` VALUES (5, 8);
INSERT INTO `blog_tags` VALUES (6, 8);
INSERT INTO `blog_tags` VALUES (7, 8);
INSERT INTO `blog_tags` VALUES (3, 9);
INSERT INTO `blog_tags` VALUES (5, 9);
INSERT INTO `blog_tags` VALUES (6, 9);
INSERT INTO `blog_tags` VALUES (7, 9);
INSERT INTO `blog_tags` VALUES (4, 10);
INSERT INTO `blog_tags` VALUES (5, 10);
INSERT INTO `blog_tags` VALUES (6, 10);
INSERT INTO `blog_tags` VALUES (7, 10);
INSERT INTO `blog_tags` VALUES (4, 11);
INSERT INTO `blog_tags` VALUES (5, 11);
INSERT INTO `blog_tags` VALUES (6, 11);
INSERT INTO `blog_tags` VALUES (5, 12);

-- ----------------------------
-- Table structure for blogs
-- ----------------------------
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `summary` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `invisibility` enum('public','private','protected') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'public',
  `views` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `likes` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blogs_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_users_blogs`(`user_id`) USING BTREE,
  CONSTRAINT `fk_users_blogs` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of blogs
-- ----------------------------
INSERT INTO `blogs` VALUES (2, '2020-10-08 14:38:57', '2020-10-12 10:44:52', NULL, '有内容的', '这是一篇有内容的博客', '# 写点代码\n```go\npackage main\nimport \"fmt\"\nfunc main() {\n  fmt.Println(\"hello, world\")\n}\n```\n- - -\n### 完毕', 'public', 40, 2, 9);
INSERT INTO `blogs` VALUES (3, '2020-10-08 19:35:06', '2020-10-11 14:54:19', '2020-10-11 22:46:34', 'ff', 'adw', 'dawdwad', 'public', 6, 0, 9);
INSERT INTO `blogs` VALUES (4, '2020-10-09 15:25:09', '2020-10-12 13:52:38', NULL, '100 个网络基础知识，看完成半个网络高手', '100 个网络基础知识，看完成半个网络高手', '## 1. 什么是链接?**\n\n链接是指两个设备之间的连接。它包括用于一个设备能够与另一个设备通信的电缆类型和协议。\n\n## 2. OSI 参考模型的层次是什么?**\n\n有 7 个 OSI 层：物理层，数据链路层，网络层，传输层，会话层，表示层和应用层。\n\n**3)什么是骨干网?**\n\n骨干网络是集中的基础设施，旨在将不同的路由和数据分发到各种网络。它还处理带宽管理和各种通道。\n\n**4)什么是 LAN?**\n\nLAN 是局域网的缩写。它是指计算机与位于小物理位置的其他网络设备之间的连接。\n\n**5)什么是节点?**\n\n节点是指连接发生的点。它可以是作为网络一部分的计算机或设备。为了形成网络连接，需要两个或更多个节点。\n\n**6)什么是路由器?**\n\n路由器可以连接两个或更多网段。这些是在其路由表中存储信息的智能网络设备，例如路径，跳数等。有了这个信息，他们就可以确定数据传输的最佳路径。路由器在 OSI 网络层运行。\n\n**7)什么是点对点链接?**\n\n它是指网络上两台计算机之间的直接连接。除了将电缆连接到两台计算机的 NIC卡之外，点对点连接不需要任何其他网络设备。\n\n**8)什么是匿名 FTP?**\n\n匿名 FTP 是授予用户访问公共服务器中的文件的一种方式。允许访问这些服务器中的数据的用户不需要识别自己，而是以匿名访客身份登录。\n\n**9)什么是子网掩码?**\n\n子网掩码与 IP 地址组合，以识别两个部分：扩展网络地址和主机地址。像 IP 地址一样，子网掩码由 32 位组成。\n\n**10)UTP 电缆允许的最大长度是多少?**\n\nUTP 电缆的单段具有 90 到 100 米的允许长度。这种限制可以通过使用中继器和开关来克服\n\n**11)什么是数据封装?**\n\n数据封装是在通过网络传输信息之前将信息分解成更小的可管理块的过程。在这个过程中，源和目标地址与奇偶校验一起附加到标题中。\n\n**12)描述网络拓扑**\n\n网络拓扑是指计算机网络的布局。它显示了设备和电缆的物理布局，以及它们如何连接到彼此。\n\n**13)什么是 VPN?**\n\nVPN 意味着虚拟专用网络，这种技术允许通过网络(如 Internet)创建安全通道。\n\n例如，VPN 允许您建立到远程服务器的安全拨号连接。\n\n**14)什么是 NAT？**\n\nNAT 是网络地址转换。这是一种协议，为公共网络上的多台计算机提供一种方式来共享到 Internet 的单一连接。\n\n**15)OSI 参考模型下网络层的工作是什么?**\n\n网络层负责数据路由，分组交换和网络拥塞控制。路由器在此层下运行。\n\n**16)网络拓扑如何影响您在建立网络时的决策?**\n\n网络拓扑决定了互连设备必须使用什么媒介。它还作为适用于设置的材料，连接器和终端的基础。\n\n**17)什么是 RIP?**\n\nRIP，路由信息协议的简称由路由器用于将数据从一个网络发送到另一个网络。\n\n它通过将其路由表广播到网络中的所有其他路由器来有效地管理路由数据。它以跳数为单位确定网络距离。\n\n**18)什么是不同的方式来保护计算机网络?**\n\n有几种方法可以做到这一点。在所有计算机上安装可靠和更新的防病毒程序。确保防火墙的设置和配置正确。用户认证也将有很大的帮助。所有这些组合将构成一个高度安全的网络。\n\n**19)什么是 NIC?**\n\nNIC 是网络接口卡(网卡)的缩写。这是连接到 PC 以连接到网络沈北。每个 NIC都有自己的 MAC 地址，用于标识网络上的 PC。\n\n**20)什么是 WAN?**\n\nWAN 代表广域网。它是地理上分散的计算机和设备的互连。它连接位于不同地区和国家/地区的网络。\n\n**21)OSI 物理层的重要性是什么?**\n\n物理层进行从数据位到电信号的转换，反之亦然。这是网络设备和电缆类型的考虑和设置。\n\n**22)TCP/IP 下有多少层?**\n\n有四层：网络层，互联网层，传输层和应用层。\n\n**23)什么是代理服务器，它们如何保护计算机网络?**\n\n代理服务器主要防止外部用户识别内部网络的 IP 地址。不知道正确的 IP 地址，甚至无法识别网络的物理位置。代理服务器可以使外部用户几乎看不到网络。\n\n**24)OSI 会话层的功能是什么?**\n\n该层为网络上的两个设备提供协议和方法，通过举行会话来相互通信。这包括设置会话，管理会话期间的信息交换以及终止会话时的解除过程。\n\n**25)实施容错系统的重要性是什么?有限吗?**\n\n容错系统确保持续的数据可用性。这是通过消除单点故障来实现的。但是，在某些情况下，这种类型的系统将无法保护数据，例如意外删除。\n\n**26)10Base-T 是什么意思?**\n\n10 是指数据传输速率，在这种情况下是 10Mbps。“Base”是指基带。T 表示双绞线，这是用于该网络的电缆。\n\n**27)什么是私有 IP 地址?**\n\n专用 IP 地址被分配用于内部网。这些地址用于内部网络，不能在外部公共网络上路由。这些确保内部网络之间不存在任何冲突，同时私有 IP 地址的范围同样可重复使用于多个内部网络，因为它们不会“看到”彼此。\n\n**28)什么是 NOS?**\n\nNOS 或网络操作系统是专门的软件，其主要任务是向计算机提供网络连接，以便能够与其他计算机和连接的设备进行通信。\n\n**29)什么是 DoS?**\n\nDoS 或拒绝服务攻击是试图阻止用户访问互联网或任何其他网络服务。这种攻击可能有不同的形式，由一群永久者组成。这样做的一个常见方法是使系统服务器过载，使其无法再处理合法流量，并将被强制重置。\n\n**30)什么是 OSI，它在电脑网络中扮演什么角色?**\n\nOSI(开放系统互连)作为数据通信的参考模型。它由 7 层组成，每层定义了网络设备如何相互连接和通信的特定方面。一层可以处理所使用的物理介质，而另一层则指示如何通过网络实际传输数据。\n\n**31)电缆被屏蔽并具有双绞线的目的是什么?**\n\n其主要目的是防止串扰。串扰是电磁干扰或噪声，可能影响通过电缆传输的数据。\n\n**32)地址共享的优点是什么?**\n\n通过使用地址转换而不是路由，地址共享提供了固有的安全性优势。这是因为互联网上的主机只能看到提供地址转换的计算机上的外部接口的公共 IP 地址，而不是内部网络上的私有 IP 地址。\n\n**33)什么是 MAC 地址?**\n\nMAC 或媒介访问控制，可以唯一地标识网络上的设备。它也被称为物理地址或以太网地址。MAC 地址由 6 个字节组成。\n\n**34)在 OSI 参考模型方面，TCP/IP 应用层的等同层或多层是什么?**\n\nTCP/IP 应用层实际上在 OSI 模型上具有三个对等体：会话层，表示层和应用层。\n\n**35)如何识别给定 IP 地址的 IP 类?**\n\n通过查看任何给定 IP 地址的第一个八位字节，您可以识别它是 A 类，B 类还是 C类。如果第一个八位字节以 0 位开头，则该地址为 Class A.如果以位 10 开头，则该地址为 B 类地址。如果从 110 开始，那么它是 C 类网络。\n\n**36)OSPF 的主要目的是什么?**\n\nOSPF 或开放最短路径优先，是使用路由表确定数据交换的最佳路径的链路状态路由协议。\n\n**37)什么是防火墙?**\n\n防火墙用于保护内部网络免受外部攻击。这些外部威胁可能是黑客谁想要窃取数据或计算机病毒，可以立即消除数据。它还可以防止来自外部网络的其他用户访问专用网络。\n\n**38)描述星形拓扑**\n\n星形拓扑由连接到节点的中央集线器组成。这是最简单的设置和维护之一。\n\n**39)什么是网关?**\n\n网关提供两个或多个网段之间的连接。它通常是运行网关软件并提供翻译服务的计算机。该翻译是允许不同系统在网络上通信的关键。\n\n**40)星型拓扑的缺点是什么?**\n\n星形拓扑的一个主要缺点是，一旦中央集线器或交换机被损坏，整个网络就变得不可用了。\n\n**41)什么是 SLIP?**\n\nSLIP 或串行线路接口协议实际上是在 UNIX 早期开发的旧协议。这是用于远程访问的协议之一。\n\n**42)给出一些私有网络地址的例子。**\n\n10.0.0.0，子网掩码为 255.0.0.0\n\n172.16.0.0，子网掩码为 255.240.0.0\n\n**43)什么是 tracert?**\n\nTracert 是一个 Windows 实用程序，可用于跟踪从路由器到目标网络的数据采集的路由。它还显示了在整个传输路由期间采用的跳数。\n\n**44)网络管理员的作用是什么?**\n\n网络管理员有许多责任，可以总结为 3 个关键功能：安装网络，配置网络设置以及网络的维护/故障排除。\n\n**45)描述对等网络的一个缺点。**\n\n当您正在访问由网络上的某个工作站共享的资源时，该工作站的性能会降低。\n\n**46)什么是混合网络?**\n\n混合网络是利用客户端 - 服务器和对等体系结构的网络设置。\n\n**47)什么是 DHCP?**\n\nDHCP 是动态主机配置协议的缩写。其主要任务是自动为网络上的设备分配 IP 地址。它首先检查任何设备尚未占用的下一个可用地址，然后将其分配给网络设备。\n\n**48)ARP 的主要工作是什么?**\n\nARP 或地址解析协议的主要任务是将已知的 IP 地址映射到 MAC 层地址。\n\n**49)什么是 TCP/IP?**\n\nTCP/IP 是传输控制协议/互联网协议的缩写。这是一组协议层，旨在在不同类型的计算机网络(也称为异构网络)上进行数据交换。\n\n**50)如何使用路由器管理网络?**\n\n路由器内置了控制台，可让您配置不同的设置，如安全和数据记录。您可以为计算机分配限制，例如允许访问的资源，或者可以浏览互联网的某一天的特定时间。您甚至可以对整个网络中看不到的网站施加限制。\n\n**51)当您希望在不同平台(如 UNIX 系统和 Windows 服务器之间)传输文件时，可以应用什么协议?**\n\n使用 FTP(文件传输协议)在这些不同的服务器之间进行文件传输。这是可能的，因为 FTP 是平台无关的。\n\n**52)默认网关的使用是什么?**\n\n默认网关提供了本地网络连接到外部网络的方法。用于连接外部网络的默认网关通常是外部路由器端口的地址。\n\n**53)保护网络的一种方法是使用密码。什么可以被认为是好的密码?**\n\n良好的密码不仅由字母组成，还包括字母和数字的组合。结合大小写字母的密码比使用所有大写字母或全部小写字母的密码有利。密码必须不能被黑客很容易猜到，比如日期，姓名，收藏夹等等。\n\n**54)UTP 电缆的正确终止率是多少?**\n\n非屏蔽双绞线网线的正常终止是 100 欧姆。\n\n**55)什么是 netstat?**\n\nNetstat 是一个命令行实用程序。它提供有关连接当前 TCP/IP 设置的有用信息。\n\n**56)C 类网络中的网络 ID 数量是多少?**\n\n对于 C 类网络，可用的网络 ID 位数为 21。可能的网络 ID 数目为 2，提高到 21或 2,097,152。每个网络 ID 的主机 ID 数量为 2，增加到 8 减去 2，或 254。\n\n**57)使用长于规定长度的电缆时会发生什么?**\n\n电缆太长会导致信号丢失。这意味着数据传输和接收将受到影响，因为信号长度下降。\n\n**58)什么常见的软件问题可能导致网络缺陷?**\n\n软件相关问题可以是以下任何一种或其组合：\n\n客户端服务器问题；应用程序冲突；配置错误；协议不匹配；安全问题；用户政策和权利问题\n\n**59)什么是 ICMP?**\n\nICMP 是 Internet 控制消息协议。它为 TCP/IP 协议栈内的协议提供消息传递和通信。这也是管理由 PING 等网络工具使用的错误信息的协议。\n\n**60)什么是 Ping?**\n\nPing 是一个实用程序，允许您检查网络上的网络设备之间的连接。您可以使用其IP 地址或设备名称(如计算机名称)ping 设备。\n\n**61)什么是点对点(P2P)?**\n\n对等是不在服务器上回复的网络。该网络上的所有 PC 都是单独的工作站。\n\n**62)什么是 DNS?**\n\nDNS 是域名系统。该网络服务的主要功能是为 TCP/IP 地址解析提供主机名。\n\n**63)光纤与其他介质有什么优势?**\n\n光纤的一个主要优点是不太容易受到电气干扰。它还支持更高的带宽，意味着可以发送和接收更多的数据。长距离信号降级也非常小。\n\n**64)集线器和交换机有什么区别?**\n\n集线器充当多端口中继器。然而，随着越来越多的设备连接到它，它将无法有效地管理通过它的流量。交换机提供了一个更好的替代方案，可以提高性能，特别是在所有端口上预期有高流量时。\n\n**65)Windows RRAS 服务支持的不同网络协议是什么?**\n\n支持三种主要的网络协议：NetBEUI，TCP/IP 和 IPX。\n\n**66)A，B 和 C 类网络中的最大网络和主机是什么?**\n\n对于 A 类，有 126 个可能的网络和 16,777,214 个主机\n\n对于 B 类，有 16,384 个可能的网络和 65,534 个主机\n\n对于 C 类，有 2,097,152 个可能的网络和 254 个主机\n\n**67)直通电缆的标准颜色顺序是什么?**\n\n橙色/白色，橙色，绿色/白色，蓝色，蓝色/白色，绿色，棕色/白色，棕色。\n\n**68)什么协议落在 TCP/IP 协议栈的应用层之下?**\n\n以下是 TCP/IP 应用层协议：FTP，TFTP，Telnet 和 SMTP。\n\n**69)您需要连接两台电脑进行文件共享。是否可以这样做，而不使用集线器或路由器?**\n\n是的，您可以使用一根电缆将两台计算机连接在一起。在这种情况下可以使用交叉型电缆。在这种设置中，一条电缆的数据传输引脚连接到另一条电缆的数据接收引脚，反之亦然。\n\n**70)什么是 ipconfig?**\n\nIpconfig 是一个常用于识别网络上计算机的地址信息的实用程序。它可以显示物理地址以及 IP 地址。\n\n**71)直通和交叉电缆有什么区别?**\n\n直通电缆用于将计算机连接到交换机，集线器或路由器。交叉电缆用于将两个类似设备连接在一起，如 PC 到 PC 或集线器到集线器。\n\n**72)什么是客户端/服务器?**\n\n客户端/服务器是一种类型的网络，其中一个或多个计算机充当服务器。服务器提供集中的资源库，如打印机和文件。客户端是指访问服务器的工作站。\n\n**73)描述网络**\n\n网络是指用于数据通信的计算机和外围设备之间的互连。可以使用有线电缆或通过无线链路进行网络连接。\n\n**74)将 NIC 卡从一台 PC 移动到另一台 PC 时，MAC 地址是否也被转移?**\n\n是的，那是因为 MAC 地址是硬连线到 NIC 电路，而不是 PC。这也意味着当 NIC卡被另一个替换时，PC 可以具有不同的 MAC 地址。\n\n**75)解释聚类支持**\n\n群集支持是指网络操作系统在容错组中连接多台服务器的能力。这样做的主要目的是在一台服务器发生故障的情况下，集群中的下一个服务器将继续进行所有处理。\n\n**76)在包含两个服务器和二十个工作站的网络中，安装防病毒程序的最佳位置是哪里?**\n\n必须在所有服务器和工作站上安装防病毒程序，以确保保护。这是因为个人用户可以访问任何工作站，并在插入可移动硬盘驱动器或闪存驱动器时引入计算机病毒。\n\n**77)描述以太网**\n\n以太网是当今使用的流行网络技术之一。它是在 20 世纪 70 年代初开发的，并且基于 IEEE 中规定的规范。以太网在局域网中使用。\n\n**78)实现环形拓扑有什么缺点?**\n\n如果网络上的一个工作站发生故障，可能会导致整个网络丢失。另一个缺点是，当需要在网络的特定部分进行调整和重新配置时，整个网络也必须被暂时关闭。\n\n**79)CSMA/CD 和 CSMA/CA 有什么区别?**\n\nCSMA/CD 或碰撞检测，每当碰撞发生时重新发送数据帧。CSMA/CA 或碰撞避免，将首先在数据传输之前广播意图发送。\n\n**80)什么是 SMTP?**\n\nSMTP 是简单邮件传输协议的缩写。该协议处理所有内部邮件，并在 TCP/IP 协议栈上提供必要的邮件传递服务。\n\n**81)什么是组播路由?**\n\n组播路由是一种有针对性的广播形式，将消息发送到所选择的用户组，而不是将其发送到子网上的所有用户。\n\n**82)加密在网络上的重要性是什么?**\n\n加密是将信息转换成用户不可读的代码的过程。然后使用秘密密钥或密码将其翻译或解密回其正常可读格式。加密有助于确保中途截获的信息仍然不可读，因为用户必须具有正确的密码或密钥。\n\n**83)如何安排和显示 IP 地址?**\n\nIP 地址显示为一系列由周期或点分隔的四位十进制数字。这种安排的另一个术语是点分十进制格式。一个例子是 192.168.101.2\n\n**84)解释认证的重要性**\n\n认证是在用户登录网络之前验证用户凭据的过程。它通常使用用户名和密码进行。这提供了限制来自网络上的有害入侵者的访问的安全手段。\n\n**85)隧道模式是什么意思?**\n\n这是一种数据交换模式，其中两个通信计算机本身不使用 IPSec。相反，将 LAN连接到中转网络的网关创建了一个使用 IPSec 协议来保护通过它的所有通信的虚拟隧道。\n\n**86)建立 WAN 链路涉及的不同技术有哪些?**\n\n模拟连接 - 使用常规电话线;数字连接 - 使用数字电话线;交换连接 - 使用发送方和接收方之间的多组链接来移动数据。\n\n**87)网格拓扑的一个优点是什么?**\n\n在一个链接失败的情况下，总会有另一个链接可用。网状拓扑实际上是最容错的网络拓扑之一。\n\n**88)在排除计算机网络问题时，可能会发生什么常见的硬件相关问题?**\n\n大部分网络由硬件组成。这些领域的问题可能包括硬盘故障，NIC 损坏甚至硬件启动。不正确的硬件配置也是其中一个疑难问题。\n\n**89)可以做什么来修复信号衰减问题?**\n\n处理这种问题的常见方法是使用中继器和集线器，因为它将有助于重新生成信号，从而防止信号丢失。检查电缆是否正确终止也是必须的。\n\n**90)动态主机配置协议如何协助网络管理?**\n\n网络管理员不必访问每台客户端计算机来配置静态 IP 地址，而是可以应用动态主机配置协议来创建称为可以动态分配给客户端的范围的 IP 地址池。\n\n**91)解释网络概念的概况?**\n\n配置文件是为每个用户设置的配置设置。例如，可以创建将用户置于组中的配置文件。\n\n**92)什么是 Sneakernet?**\n\nSneakernet 被认为是最早的联网形式，其中使用可移动介质(如磁盘，磁带)物理传输数据。\n\n**93)IEEE 在计算机网络中的作用是什么?**\n\nIEEE 或电气和电子工程师学会是由电气和电子设备标准发布和管理的工程师组成的组织。这包括网络设备，网络接口，cablings 和连接器。\n\n**94)TCP/IP Internet 层下有哪些协议?**\n\n该层管理的协议有 4 种。这些是 ICMP，IGMP，IP 和 ARP。\n\n**95)谈到网络，什么是权限?**\n\n权限是指在网络上执行特定操作的授权许可。网络上的每个用户可以分配个人权限，具体取决于该用户必须允许的内容。\n\n**96)建立 VLAN 的一个基本要求是什么?**\n\n需要一个 VLAN，因为在交换机级别只有一个广播域，这意味着每当新用户连接时，该信息都会传播到整个网络。交换机上的 VLAN 有助于在交换机级别创建单独的广播域。它用于安全目的。\n\n**97)什么是 IPv6?**\n\nIPv6 或 Internet 协议版本 6 被开发以替代 IPv4。目前，IPv4 正在用于控制互联网流量，但 IPv4 已经饱和。IPv6 能够克服这个限制。\n\n**98)什么是 RSA 算法?**\n\nRSA 是 Rivest-Shamir-Adleman 算法的缩写。它是目前最常用的公钥加密算法。\n\n**99)什么是网格拓扑?**\n\n网格拓扑是一种设置，其中每个设备都直接连接到网络上的每个其他设备。因此，它要求每个设备具有至少两个网络连接。\n\n**100)100Base-FX 网络的最大段长度是多少?**\n\n使用 100Base-FX 的网段的最大允许长度为 412 米。整个网络的最大长度为 5 公里。', 'public', 94, 4, 9);
INSERT INTO `blogs` VALUES (5, '2020-10-09 21:19:19', '2020-10-12 10:28:51', NULL, 'go从入门到卸载', '本文仅用两分钟带大家学习golang的入门到卸载', '### 1. 安装\n\n##### 去Golang官网下载安装最新版本\n\n### 2. Hello, World\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, world!\")\n}\n```\n\n### 3. 卸载\n\n好了，现在可以卸载了', 'public', 34, 1, 9);
INSERT INTO `blogs` VALUES (6, '2020-10-10 17:36:59', '2020-10-12 13:52:19', NULL, '顶起顶起', '地区的去访问是否为分为氛围', '### 1. 封杀封杀\n#### 1.1 官方热歌\n#### 1.2 丰富\n### 2. 发生过担任韩国\n#### 个废人个人\n### 3. 风格然后他\n#### `sudo apt update`\n### 纷纷为', 'public', 30, 2, 9);
INSERT INTO `blogs` VALUES (7, '2020-10-11 21:11:31', '2020-10-12 10:44:55', NULL, '发顺丰', '随便写一篇', '# 大标题\n## 小标题\n###\n- - -\n```go\nfmt.Println(\"What)\n```\n- - -', 'public', 12, 5, 2);

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `readed` tinyint(1) NOT NULL DEFAULT 0,
  `disabled` tinyint(1) NOT NULL DEFAULT 0,
  `user_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'anonymous',
  `user_avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `blog_id` bigint(20) UNSIGNED NOT NULL,
  `invisibility` enum('public','private') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'public',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comments_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_blogs_comments`(`blog_id`) USING BTREE,
  INDEX `fk_users_comments`(`user_id`) USING BTREE,
  CONSTRAINT `fk_blogs_comments` FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '2020-10-09 16:17:47', '2020-10-09 16:17:47', NULL, '嗯，这篇博客写得很详细', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (2, '2020-10-09 16:18:50', '2020-10-09 16:18:50', NULL, '嗯，这篇博客写得很详细x2', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (3, '2020-10-09 16:18:53', '2020-10-09 16:18:53', NULL, '嗯，这篇博客写得很详细x3', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (4, '2020-10-09 16:18:56', '2020-10-09 16:18:56', NULL, '嗯，这篇博客写得很详细x4', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (5, '2020-10-09 16:19:00', '2020-10-09 16:19:00', NULL, '嗯，这篇博客写得很详细x5', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (6, '2020-10-09 16:44:55', '2020-10-09 16:44:55', NULL, '嗯，这篇博客写得很详细x5', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (7, '2020-10-09 16:45:05', '2020-10-09 16:45:05', NULL, '我也觉得', 0, 0, 'Chika', 'http://localhost:8000/api/v1/data/images/a2fb0cf066816f964f6d6d2a7fba68d8.png', 9, 4, 'public');
INSERT INTO `comments` VALUES (8, '2020-10-09 16:59:24', '2020-10-09 16:59:24', NULL, '啦啦啦', 0, 0, 'fasico', 'http://localhost:8000/api/v1/data/images/731c44d50a7d04e381db294b8f78c18e.jpg', 9, 4, 'public');
INSERT INTO `comments` VALUES (9, '2020-10-09 16:59:54', '2020-10-09 16:59:54', NULL, '啦啦啦', 0, 0, 'fasico', 'http://localhost:8000/api/v1/data/images/731c44d50a7d04e381db294b8f78c18e.jpg', 9, 4, 'public');
INSERT INTO `comments` VALUES (10, '2020-10-09 17:00:03', '2020-10-09 17:00:03', NULL, 'ggg ', 0, 0, 'fasico', 'http://localhost:8000/api/v1/data/images/731c44d50a7d04e381db294b8f78c18e.jpg', 9, 4, 'public');
INSERT INTO `comments` VALUES (11, '2020-10-09 17:00:13', '2020-10-09 17:00:13', NULL, '可以了吗', 0, 0, 'fasico', 'http://localhost:8000/api/v1/data/images/731c44d50a7d04e381db294b8f78c18e.jpg', 9, 4, 'public');
INSERT INTO `comments` VALUES (12, '2020-10-09 17:01:42', '2020-10-09 17:01:42', NULL, '就这？', 0, 0, 'fasico', 'http://localhost:8000/api/v1/data/images/731c44d50a7d04e381db294b8f78c18e.jpg', 9, 2, 'public');
INSERT INTO `comments` VALUES (13, '2020-10-09 18:21:41', '2020-10-09 18:21:41', NULL, 'what', 0, 0, 'test', '', 9, 4, 'public');
INSERT INTO `comments` VALUES (14, '2020-10-09 18:24:32', '2020-10-09 18:24:32', NULL, 'what?', 0, 0, 'admin', '', 9, 4, 'public');
INSERT INTO `comments` VALUES (15, '2020-10-09 21:19:50', '2020-10-09 21:19:50', NULL, '太厉害了', 0, 0, 'lalala', '', 9, 5, 'public');
INSERT INTO `comments` VALUES (16, '2020-10-10 13:32:04', '2020-10-10 13:32:04', NULL, '密码', 0, 0, 'aaaa', '', 9, 2, 'public');
INSERT INTO `comments` VALUES (17, '2020-10-11 18:49:54', '2020-10-11 18:49:54', NULL, 'damn it\n', 0, 0, 'admin', '', 9, 6, 'public');
INSERT INTO `comments` VALUES (18, '2020-10-12 10:45:35', '2020-10-12 10:45:35', NULL, '是的归属感', 0, 0, 'yarn', 'http://localhost:8000/api/v1/data/images/ee11e4095645325071b979b4186c9b8d.jpg', 9, 7, 'public');

-- ----------------------------
-- Table structure for diaries
-- ----------------------------
DROP TABLE IF EXISTS `diaries`;
CREATE TABLE `diaries`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `invisibility` enum('public','private','protected') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'private',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_diaries_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_users_diaries`(`user_id`) USING BTREE,
  CONSTRAINT `fk_users_diaries` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of diaries
-- ----------------------------
INSERT INTO `diaries` VALUES (1, '2020-10-07 23:50:26', '2020-10-07 23:50:26', '2020-10-11 22:46:03', '完成了完成了，睡觉', 'public', 9);
INSERT INTO `diaries` VALUES (2, '2020-10-08 16:16:27', '2020-10-08 16:16:27', '2020-10-11 22:46:01', '吃饭', 'public', 9);
INSERT INTO `diaries` VALUES (3, '2020-10-08 16:25:30', '2020-10-08 16:25:30', NULL, '私密的日记', 'private', 9);
INSERT INTO `diaries` VALUES (4, '2020-10-08 16:28:01', '2020-10-08 16:28:01', NULL, '纷纷为威风威风', 'public', 9);
INSERT INTO `diaries` VALUES (5, '2020-10-08 16:28:32', '2020-10-08 16:28:32', '2020-10-11 18:31:03', 'cxwcw ', 'public', 9);
INSERT INTO `diaries` VALUES (6, '2020-10-08 16:28:40', '2020-10-08 16:28:40', '2020-10-11 18:30:59', '微风威锋网', 'public', 9);
INSERT INTO `diaries` VALUES (7, '2020-10-08 16:28:46', '2020-10-08 16:28:46', '2020-10-11 18:31:00', '潜伏期无法', 'private', 9);
INSERT INTO `diaries` VALUES (8, '2020-10-08 19:34:11', '2020-10-08 19:34:11', '2020-10-11 18:30:55', 'gfgg', 'private', 9);
INSERT INTO `diaries` VALUES (9, '2020-10-09 18:25:02', '2020-10-09 18:25:02', NULL, '淦，终于写得差不多了', 'public', 9);
INSERT INTO `diaries` VALUES (11, '2020-10-11 21:03:23', '2020-10-11 21:03:23', NULL, 'rhrhrth二人哥哥俄国', 'public', 9);
INSERT INTO `diaries` VALUES (12, '2020-10-11 21:09:38', '2020-10-11 21:09:38', NULL, '变小了变小了', 'public', 9);

-- ----------------------------
-- Table structure for diary_tags
-- ----------------------------
DROP TABLE IF EXISTS `diary_tags`;
CREATE TABLE `diary_tags`  (
  `diary_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `tag_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`diary_id`, `tag_id`) USING BTREE,
  INDEX `fk_diary_tags_tag`(`tag_id`) USING BTREE,
  CONSTRAINT `fk_diary_tags_diary` FOREIGN KEY (`diary_id`) REFERENCES `diaries` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_diary_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of diary_tags
-- ----------------------------
INSERT INTO `diary_tags` VALUES (9, 1);
INSERT INTO `diary_tags` VALUES (12, 1);
INSERT INTO `diary_tags` VALUES (3, 2);
INSERT INTO `diary_tags` VALUES (9, 2);
INSERT INTO `diary_tags` VALUES (12, 2);
INSERT INTO `diary_tags` VALUES (2, 3);
INSERT INTO `diary_tags` VALUES (3, 3);
INSERT INTO `diary_tags` VALUES (7, 3);
INSERT INTO `diary_tags` VALUES (9, 3);
INSERT INTO `diary_tags` VALUES (11, 3);
INSERT INTO `diary_tags` VALUES (12, 3);
INSERT INTO `diary_tags` VALUES (2, 4);
INSERT INTO `diary_tags` VALUES (3, 4);
INSERT INTO `diary_tags` VALUES (6, 4);
INSERT INTO `diary_tags` VALUES (7, 4);
INSERT INTO `diary_tags` VALUES (11, 4);
INSERT INTO `diary_tags` VALUES (12, 4);
INSERT INTO `diary_tags` VALUES (2, 5);
INSERT INTO `diary_tags` VALUES (4, 5);
INSERT INTO `diary_tags` VALUES (6, 5);
INSERT INTO `diary_tags` VALUES (7, 5);
INSERT INTO `diary_tags` VALUES (8, 5);
INSERT INTO `diary_tags` VALUES (11, 5);
INSERT INTO `diary_tags` VALUES (12, 5);
INSERT INTO `diary_tags` VALUES (5, 6);
INSERT INTO `diary_tags` VALUES (8, 6);
INSERT INTO `diary_tags` VALUES (12, 6);

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `text` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `color` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `use_for` enum('diary','blog') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'diary',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_tags_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tags
-- ----------------------------
INSERT INTO `tags` VALUES (1, '2020-09-23 19:03:23', '2020-09-29 14:51:19', NULL, 'mdi-cup', '奶茶', 'brown darken-3', 'diary');
INSERT INTO `tags` VALUES (2, '2020-09-23 19:04:38', '2020-09-29 16:25:34', NULL, 'mdi-heart-outline', '开心', 'red', 'diary');
INSERT INTO `tags` VALUES (3, '2020-09-23 19:05:20', '2020-09-29 16:25:34', NULL, 'mdi-emoticon-sad-outline', '难过', 'grey', 'diary');
INSERT INTO `tags` VALUES (4, '2020-09-23 19:06:04', '2020-09-29 16:25:34', NULL, 'mdi-movie-outline', '电影', 'orange', 'diary');
INSERT INTO `tags` VALUES (5, '2020-09-23 19:06:34', '2020-09-29 16:25:34', NULL, 'mdi-white-balance-sunny', '晴朗', 'teal', 'diary');
INSERT INTO `tags` VALUES (6, '2020-09-23 19:14:17', '2020-09-29 14:51:18', NULL, 'mdi-weather-pouring', '大雨', 'blue-grey', 'diary');
INSERT INTO `tags` VALUES (7, '2020-09-29 15:01:26', '2020-09-29 15:32:38', NULL, 'mdi-label', 'Vue', 'green', 'blog');
INSERT INTO `tags` VALUES (8, '2020-09-29 15:01:44', '2020-09-29 15:34:16', NULL, 'mdi-label', 'Go', 'teal', 'blog');
INSERT INTO `tags` VALUES (9, '2020-09-29 15:01:52', '2020-09-29 15:56:09', NULL, 'mdi-label', 'python', 'blue', 'blog');
INSERT INTO `tags` VALUES (10, '2020-09-29 15:31:39', '2020-09-29 15:56:09', NULL, 'mdi-label', 'java', 'red', 'blog');
INSERT INTO `tags` VALUES (11, '2020-09-29 15:31:57', '2020-09-29 15:34:16', NULL, 'mdi-label', 'vuetify', 'blue', 'blog');
INSERT INTO `tags` VALUES (12, '2020-09-29 15:32:10', '2020-09-29 15:34:38', NULL, 'mdi-label', 'bootstrap', 'purple', 'blog');

-- ----------------------------
-- Table structure for todos
-- ----------------------------
DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `text` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `done` tinyint(1) NULL DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_todos_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_users_to_dos`(`user_id`) USING BTREE,
  CONSTRAINT `fk_users_to_dos` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of todos
-- ----------------------------
INSERT INTO `todos` VALUES (1, '2020-10-07 23:50:05', '2020-10-12 14:51:36', NULL, '睡觉', 1, 9);
INSERT INTO `todos` VALUES (2, '2020-10-08 19:33:44', '2020-10-08 19:33:59', '2020-10-12 14:51:39', 'ff', 1, 9);
INSERT INTO `todos` VALUES (3, '2020-10-08 22:15:26', '2020-10-09 18:25:28', '2020-10-09 18:25:29', '啦啦啦', 1, 9);
INSERT INTO `todos` VALUES (4, '2020-10-09 18:25:24', '2020-10-09 18:25:24', NULL, '吃饭', 0, 9);
INSERT INTO `todos` VALUES (6, '2020-10-11 21:07:12', '2020-10-11 21:07:21', NULL, '吃饭', 1, 9);
INSERT INTO `todos` VALUES (7, '2020-10-11 22:41:01', '2020-10-11 22:41:01', NULL, 'chifan', 0, 9);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role` enum('user','admin') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'user',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (2, '2020-10-08 16:22:46', '2020-10-11 21:13:56', NULL, 'test', '$04$vXSj4VUdxxyjaEsf9tSWA.7V3X7M4si22fs7WLMuDOk0Cy2bYEF4S', 'user', 'fasico@qq.com', '', 'http://localhost:8000/api/v1/data/images/386075b3c0d73ce1c0862ee41b4e2938.jpg', 'simple, cool, fast');
INSERT INTO `users` VALUES (3, '2020-10-08 19:33:21', '2020-10-08 19:33:21', NULL, 'aaaa', '$2a$04$vXSj4VUdxxyjaEsf9tSWA.7V3X7M4si22fs7WLMuDOk0Cy2bYEF4S', 'user', '', '', '', '');
INSERT INTO `users` VALUES (8, '2020-10-11 21:26:24', '2020-10-11 21:26:24', NULL, 'aaaa', '$04$8SF71rZI9SxMx/XWNNflbeytJTK6h.RQFiJuze7tHlhvcaShrGxTy', 'user', '', '', 'http://localhost:8000/api/v1/data/images/be665a1ae112183a73e8db209f416c3d.png', '');
INSERT INTO `users` VALUES (9, '2020-10-11 21:32:22', '2020-10-11 21:33:32', NULL, 'yarn', '$2a$04$8SF71rZI9SxMx/XWNNflbeytJTK6h.RQFiJuze7tHlhvcaShrGxTy', 'user', '', '13340257474', 'http://localhost:8000/api/v1/data/images/ee11e4095645325071b979b4186c9b8d.jpg', '');

SET FOREIGN_KEY_CHECKS = 1;
