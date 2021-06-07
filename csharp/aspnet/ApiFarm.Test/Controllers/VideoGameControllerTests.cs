using System.Collections.Generic;
using ApiFarm.Controllers;
using ApiFarm.Models.Impl;
using ApiFarm.Services;
using ApiFarm.Test.Helpers;
using ApiFarm.Utils;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Controllers
{
    public class VideoGameControllerTests
    {
        private VideoGamesController subject;

        private Mock<IService<VideoGame>> mockVideoGameService;

        [SetUp]
        protected void SetUp()
        {
            mockVideoGameService = new Mock<IService<VideoGame>>();

            subject = new VideoGamesController(mockVideoGameService.Object);
        }

        private class GetShould : VideoGameControllerTests
        {
            [Test]
            public void ReturnOkResponse()
            {
                // Arrange
                var strId = "12";
                var id = 12u;
                var storedVideoGame = new VideoGame();
                var mockQuery = new Mock<IQuery<VideoGame>>();

                mockVideoGameService.Setup(m => m.Get(id)).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Result).Returns(storedVideoGame);

                // Act
                var objectResult = subject.Get(strId);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status200OK);
                objectResult.Value.ShouldBe(storedVideoGame);
            }

            [Test]
            public void ReturnBadRequestResponse()
            {
                // Arrange
                var strId = "invalid!";

                // Act
                var objectResult = subject.Get(strId);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status400BadRequest);
                objectResult.Value.ShouldBe(ResponseMessages.Id.IsInvalid(strId));
            }

            [Test]
            public void ReturnNotFoundResponse()
            {
                // Arrange
                var strId = "99";
                var id = 99u;
                var mockQuery = new Mock<IQuery<VideoGame>>();
                var queryMessage = "NOT FOUND!";

                mockVideoGameService.Setup(m => m.Get(id)).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Code).Returns(404);
                mockQuery.Setup(m => m.Message).Returns(queryMessage);

                // Act
                var objectResult = subject.Get(strId);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status404NotFound);
                objectResult.Value.ShouldBe(queryMessage);
            }
        }

        private class GetAllShould : VideoGameControllerTests
        {
            [Test]
            public void ReturnOkResponse()
            {
                // Arrange
                var storedVideoGames = new Mock<IEnumerable<VideoGame>>();
                var mockQuery = new Mock<IQuery<IEnumerable<VideoGame>>>();

                mockVideoGameService.Setup(m => m.GetAll()).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Result).Returns(storedVideoGames.Object);

                // Act
                var objectResult = subject.GetAll().AsObjectResult();

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status200OK);
                objectResult.Value.ShouldBe(storedVideoGames.Object);
            }
        }

        private class AddShould : VideoGameControllerTests
        {
            [Test]
            public void ReturnCreatedResponse()
            {
                // Arrange
                var videoGame = new VideoGame();
                var storedVideoGame = new VideoGame { Id = 5 };
                var mockQuery = new Mock<IQuery<VideoGame>>();

                mockVideoGameService.Setup(m => m.Add(videoGame)).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Result).Returns(storedVideoGame);

                // Act
                var objectResult = subject.Post(videoGame);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status201Created);
                objectResult.Value.ShouldBe(storedVideoGame);
                objectResult.ShouldBeOfType<CreatedResult>();
                (objectResult as CreatedResult).Location.ShouldBe("video_games/5");
            }

            [Test]
            public void ReturnBadRequestResponse()
            {
                // Arrange
                var videoGame = new VideoGame();
                var mockQuery = new Mock<IQuery<VideoGame>>();
                var queryMessage = "FAILED!";

                mockVideoGameService.Setup(m => m.Add(videoGame)).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Code).Returns(400);
                mockQuery.Setup(m => m.Message).Returns(queryMessage);

                // Act
                var objectResult = subject.Post(videoGame);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status400BadRequest);
                objectResult.Value.ShouldBe(queryMessage);
            }
        }

        private class UpdateShould : VideoGameControllerTests
        {
            [Test]
            public void ReturnOkResponse()
            {
                // Arrange
                var strId = "5";
                var videoGameUpdateValues = new VideoGame();
                var mockQuery = new Mock<IQuery<VideoGame>>();
                var updatedVideoGame = new VideoGame();

                mockVideoGameService.Setup(m => m.Update(5, videoGameUpdateValues)).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Result).Returns(updatedVideoGame);

                // Act
                var objectResult = subject.Put(strId, videoGameUpdateValues);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status200OK);
                objectResult.Value.ShouldBe(updatedVideoGame);
            }

            [Test]
            public void ReturnBadRequestResponse()
            {
                // Arrange
                var strId = "invalid!";

                // Act
                var objectResult = subject.Put(strId, new VideoGame());

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status400BadRequest);
                objectResult.Value.ShouldBe(ResponseMessages.Id.IsInvalid(strId));
            }

            [Test]
            public void ReturnNotFoundResponse()
            {
                // Arrange
                var strId = "99";
                var videoGameUpdateValues = new VideoGame();
                var mockQuery = new Mock<IQuery<VideoGame>>();
                var queryMessage = "NOT FOUND!";

                mockVideoGameService.Setup(m => m.Update(99, videoGameUpdateValues)).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Code).Returns(404);
                mockQuery.Setup(m => m.Message).Returns(queryMessage);

                // Act
                var objectResult = subject.Put(strId, videoGameUpdateValues);

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status404NotFound);
                objectResult.Value.ShouldBe(queryMessage);
            }
        }
    }
}
