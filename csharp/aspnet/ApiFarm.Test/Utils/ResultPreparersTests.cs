using System.Collections.Generic;
using ApiFarm.Models.Impl;
using NUnit.Framework;
using Shouldly;
using static ApiFarm.Utils.ResultPreparers;

namespace ApiFarm.Test.Utils
{
    public class ResultPreparersTests
    {
        public class PrepVideoGameShould : ResultPreparersTests
        {
            [Test]
            public void NotSetEmptyListWhenValued()
            {
                // Arrange
                var developers = new List<string>();
                var publishers = new List<string>() { "A" };
                var directors = new List<string>() { "A", "B" };
                var producers = new List<string>() { "A", "B", "C" };
                var designers = new List<string>() { "A", "B" };
                var programmers = new List<string>() { "A", "B", "C" };
                var artists = new List<string>() { "A", "B", "C" };
                var composers = new List<string>() { "A", "B" };
                var platforms = new List<string>();

                var videoGame = new VideoGame
                {
                    Developers = developers,
                    Publishers = publishers,
                    Directors = directors,
                    Producers = producers,
                    Designers = designers,
                    Programmers = programmers,
                    Artists = artists,
                    Composers = composers,
                    Platforms = platforms,
                };

                // Act
                PrepVideoGame(videoGame);

                // Assert
                videoGame.Developers.ShouldBe(developers);
                videoGame.Publishers.ShouldBe(publishers);
                videoGame.Directors.ShouldBe(directors);
                videoGame.Producers.ShouldBe(producers);
                videoGame.Designers.ShouldBe(designers);
                videoGame.Programmers.ShouldBe(programmers);
                videoGame.Artists.ShouldBe(artists);
                videoGame.Composers.ShouldBe(composers);
                videoGame.Platforms.ShouldBe(platforms);
            }

            [Test]
            public void SetEmptyListsWhenNull()
            {
                // Arrange
                var developers = default(List<string>);
                var publishers = default(List<string>);
                var directors = default(List<string>);
                var producers = default(List<string>);
                var designers = default(List<string>);
                var programmers = default(List<string>);
                var artists = default(List<string>);
                var composers = default(List<string>);
                var platforms = default(List<string>);

                var videoGame = new VideoGame
                {
                    Developers = developers,
                    Publishers = publishers,
                    Directors = directors,
                    Producers = producers,
                    Designers = designers,
                    Programmers = programmers,
                    Artists = artists,
                    Composers = composers,
                    Platforms = platforms,
                };

                // Act
                PrepVideoGame(videoGame);

                // Assert
                videoGame.Developers.ShouldNotBe(developers);
                videoGame.Publishers.ShouldNotBe(publishers);
                videoGame.Directors.ShouldNotBe(directors);
                videoGame.Producers.ShouldNotBe(producers);
                videoGame.Designers.ShouldNotBe(designers);
                videoGame.Programmers.ShouldNotBe(programmers);
                videoGame.Artists.ShouldNotBe(artists);
                videoGame.Composers.ShouldNotBe(composers);
                videoGame.Platforms.ShouldNotBe(platforms);
            }
        }
    }
}
