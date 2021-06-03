using System.Collections.Generic;
using ApiFarm.Models;
using ApiFarm.Repositories;
using ApiFarm.Utils;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Repositories
{
    public class InMemoryTests
    {
        private InMemory<Model> subject;

        private Mock<ICloner<Model>> mockCloner;
        private List<Model> models;

        [SetUp]
        public void Setup()
        {
            mockCloner = new Mock<ICloner<Model>>();
            models = new List<Model>();
            subject = new InMemory<Model>(mockCloner.Object, models);
        }

        public class Model : IModel
        {
            public uint Id { get; set; }
        }

        private class ResetShould : InMemoryTests
        {
            [Test]
            public void ClearModelsList()
            {
                // Arrange
                models.Add(new Model());
                models.Add(new Model());
                models.Add(new Model());

                // Act
                subject.Reset();

                // Assert
                models.ShouldBeEmpty();
            }
        }

        private class GetShould : InMemoryTests
        {
            [Test]
            public void ReturnCloneOfModelWithId()
            {
                // Arrange
                var id = 5u;
                var desiredModel = new Model { Id = id };
                models.Add(new Model { Id = 2 });
                models.Add(desiredModel);
                models.Add(new Model { Id = 9 });
                var expected = new Model();

                mockCloner.Setup(m => m.Clone(desiredModel)).Returns(expected);

                // Act
                var actual = subject.Get(id);

                // Assert
                actual.ShouldBe(expected);
            }

            [Test]
            public void ReturnDefaultWhenNoModelWithIdExists()
            {
                // Act
                var actual = subject.Get(5);

                // Assert
                actual.ShouldBeNull();
            }
        }

        private class GetAllShould : InMemoryTests
        {
            [Test]
            public void ReturnEmptyModelsList()
            {
                // Arrange
                subject = new InMemory<Model>(mockCloner.Object);

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldBeEmpty();
            }

            [Test]
            public void ReturnListOfClonedModels()
            {
                // Arrange
                var originalModel = new Model();
                models.Add(originalModel);
                var cloneModel = new Model();

                mockCloner.Setup(m => m.Clone(originalModel)).Returns(cloneModel);

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldNotBeSameAs(models);
                retrievedModels.ShouldContain(cloneModel);
            }
        }

        private class AddShould : InMemoryTests
        {
            [Test]
            public void SetIncrementedIdOfModelAndAddToList()
            {
                // Arrange
                var originalModel1 = new Model();
                var cloneModel1 = new Model();
                var originalModel2 = new Model();
                var cloneModel2 = new Model();

                mockCloner.Setup(m => m.Clone(originalModel1)).Returns(cloneModel1);
                mockCloner.Setup(m => m.Clone(originalModel2)).Returns(cloneModel2);

                // Act
                var addedModel1 = subject.Add(originalModel1);
                var addedModel2 = subject.Add(originalModel2);

                // Assert
                models.ShouldContain(cloneModel1);
                models.ShouldContain(cloneModel2);

                addedModel1.ShouldBe(originalModel1);
                addedModel2.ShouldBe(originalModel2);

                addedModel1.Id.ShouldBe(1u);
                addedModel2.Id.ShouldBe(2u);
            }
        }

        private class UpdateShould : InMemoryTests
        {
            [Test]
            public void ReplaceIdentifiedModelWithClonedUpdatedVersion()
            {
                // Arrange
                var modelToUpdate = new Model { Id = 3 };
                models.Add(new Model { Id = 1 });
                models.Add(modelToUpdate);
                models.Add(new Model { Id = 9 });
                var updatedModel = new Model { Id = 3 };
                var updatedModelToStore = new Model();

                mockCloner.Setup(m => m.Clone(updatedModel)).Returns(updatedModelToStore);

                // Act
                var actual = subject.Update(updatedModel);

                // Assert
                actual.ShouldBe(updatedModel);

                models.ShouldNotContain(modelToUpdate);
                models.ShouldContain(updatedModelToStore);
            }

            [Test]
            public void ReturnNullWhenNoIdentifiedInstanceToUpdateCouldBeFound()
            {
                // Arrange
                var updatedModel = new Model { Id = 5 };

                // Act
                var actual = subject.Update(updatedModel);

                // Assert
                actual.ShouldBeNull();

                mockCloner.Verify(m => m.Clone(It.IsAny<Model>()), Times.Never);
            }
        }

        private class DeleteShould : InMemoryTests
        {
            [Test]
            public void ReturnModelWhichWasRemovedFromList()
            {
                // Arrange
                var id = 7u;
                var modelToDelete = new Model { Id = id };
                models.Add(new Model { Id = 1 });
                models.Add(modelToDelete);
                models.Add(new Model { Id = 15 });

                // Act
                var deletedModel = subject.Delete(id);

                // Assert
                deletedModel.ShouldBe(modelToDelete);
                models.ShouldNotContain(modelToDelete);
                models.Count.ShouldBe(2);
            }

            [Test]
            public void ReturnNullWhenModelWithIdentifierNotFound()
            {
                // Act
                var deletedModel = subject.Delete(99u);

                // Assert
                deletedModel.ShouldBeNull();
            }
        }
    }
}
