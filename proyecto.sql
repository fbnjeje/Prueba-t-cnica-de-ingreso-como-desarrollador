-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 24-10-2023 a las 23:46:04
-- Versión del servidor: 10.4.27-MariaDB
-- Versión de PHP: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `proyecto`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `datosvideo`
--

CREATE TABLE `datosvideo` (
  `ID` int(11) NOT NULL,
  `Tema` varchar(50) NOT NULL,
  `Creador` varchar(50) NOT NULL,
  `Descripcion` varchar(250) NOT NULL,
  `Titulo` varchar(50) NOT NULL,
  `Estado` tinyint(1) NOT NULL,
  `CreatedOn` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `datosvideo`
--

INSERT INTO `datosvideo` (`ID`, `Tema`, `Creador`, `Descripcion`, `Titulo`, `Estado`, `CreatedOn`) VALUES
(1, '[value-2]', 'Musica', 'Aprende mas en nuestra pagina Web', 'Musica para todos', 0, '0000-00-00');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
